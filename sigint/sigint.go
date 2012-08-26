package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("Started")

	go func() {
		sigchan := make(chan os.Signal, 10)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		fmt.Println("\nCaught interrupt.")
		// TODO: Last actions and wait for all operations to end.
		fmt.Println("Finished")
		os.Exit(0)
	}()

	for {
		fmt.Println("Sleeping...")
		time.Sleep(1 * time.Second)
	}
}
