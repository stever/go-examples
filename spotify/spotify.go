package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
	"unsafe"
)

/*
#include <stdlib.h>
extern void cspotify_login(char *username, char *password);
extern void cspotify_logged_in();
extern void cspotify_notify_main_thread();
extern int cspotify_music_delivery();
extern void cspotify_metadata_updated();
extern void cspotify_play_token_lost();
extern void cspotify_end_of_track();
#cgo LDFLAGS: -llibspotify
*/
import "C"

func main() {
	fmt.Println("Started")

	go func() {
		sigchan := make(chan os.Signal, 10)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		fmt.Println("Program killed !")

		// do last actions and wait for all write operations to end

		os.Exit(0)
	}()

	// Username and password.
	var username *C.char = C.CString("username")
	var password *C.char = C.CString("password")
	defer C.free(unsafe.Pointer(username))
	defer C.free(unsafe.Pointer(password))

	C.cspotify_login(username, password)

	for {
		fmt.Println("Sleeping...")
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Finished")
}

//export cspotify_logged_in
func cspotify_logged_in() {
	fmt.Println("logged_in")
}

//export cspotify_notify_main_thread
func cspotify_notify_main_thread() {
	fmt.Println("notify_main_thread")
}

//export cspotify_music_delivery
func cspotify_music_delivery() C.int {
	fmt.Println("music_delivery")
	return 0
}

//export cspotify_metadata_updated
func cspotify_metadata_updated() {
	fmt.Println("metadata_updated")
}

//export cspotify_play_token_lost
func cspotify_play_token_lost() {
	fmt.Println("play_token_lost")
}

//export cspotify_end_of_track
func cspotify_end_of_track() {
	fmt.Println("end_of_track")
}
