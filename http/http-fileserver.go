package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
)

func server(port int) {
	mime.AddExtensionType(".svg", "image/svg+xml")
	htdocs, _ := filepath.Abs("htdocs")
	http.Handle("/", http.FileServer(http.Dir(htdocs)))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func main() {
	go func() {
		sigchan := make(chan os.Signal, 10)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		fmt.Println("\nCaught interrupt.")

		// Shutdown
		log.Println("Finished")
		os.Exit(0)
	}()

	// Start processes
	port := 8080
	go server(port)
	log.Printf("Server started on http://localhost:%d", port)

	// Sleep forever
	select {}
}
