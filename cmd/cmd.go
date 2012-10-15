package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("date").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("The date is", string(out))
}
