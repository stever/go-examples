// This is an example of running a command and capturing its output.

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {

	// Example command to call.
	cmd := exec.Command("ls", "-al")

	// We'll write the output of the command to the following byte buffer.
	var out bytes.Buffer

	// Get stdout reader.
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Get stderr reader.
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Start the command.
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Pipe any error output to stderr of this program.
	go io.Copy(os.Stderr, stderr)

	// Copy output from the command to a buffer.
	// This is implemented as a function as it would be useful to process the command output.
	// If just streaming the output you could use "go io.Copy(&out, stdout)" instead.
	go func(src io.ReadCloser, dst *bytes.Buffer) {
		var written int64
		var ErrShortWrite = errors.New("short write")
		var EOF = errors.New("EOF")
		buf := make([]byte, 32*1024)
		for {
			nr, er := src.Read(buf)
			if nr > 0 {
				nw, ew := dst.Write(buf[0:nr])
				if nw > 0 {
					written += int64(nw)
				}
				if ew != nil {
					err = ew
					break
				}
				if nr != nw {
					err = ErrShortWrite
					break
				}
			}
			if er == EOF {
				break
			}
			if er != nil {
				err = er
				break
			}
		}
	}(stdout, &out)

	cmd.Wait()
	fmt.Print(out.String())
}
