// Logging example.
// http://play.golang.org/p/QFheQeChIn
// https://groups.google.com/d/topic/golang-nuts/gU7oQGoCkmg/discussion

package main

import "log"

const debug debugging = true // or flip to false

type debugging bool

func (d debugging) Printf(format string, args ...interface{}) {
	if d {
		log.Printf(format, args...)
	}
}

func main() {
	debug.Printf("foo %d %.2f", 42, 12.7)
}
