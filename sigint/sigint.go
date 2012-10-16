package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("Started")

	// Interrupt handler.
	// NOTE: This doesn't currently work on Windows (2012-10-16 go1.0.3)
	go func() {
		sigchan := make(chan os.Signal, 10)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		fmt.Println("\nCaught interrupt.")
		// NOTE: Do last actions here and wait for all operations to end.
		fmt.Println("Finished")
		os.Exit(0)
	}()

	for {
		fmt.Println("Sleeping...")
		time.Sleep(1 * time.Second)
	}
}
