// This is an example using JSON to support something like a Playdar resolver.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type settings struct {
	Name       string `json:"name"`
	Weight     int    `json:"weight"`
	TargetTime int    `json:"targettime"`
}

type query struct {
	Qid    string `json:"qid"`
	Artist string `json:"artist"`
	Track  string `json:"track"`
	Title  string `json:"title"`
}

func msgtype(msg []byte, msgtype string) ([]byte, error) {
	var f interface{}
	err := json.Unmarshal(msg, &f)
	if err != nil {
		log.Fatal(err)
	}
	m := f.(map[string]interface{})
	m["_msgtype"] = msgtype
	return json.Marshal(m)
}

func msettings(inf settings) ([]byte, error) {
	bytes, err := json.Marshal(inf)
	if err != nil {
		log.Fatal(err)
	}
	return msgtype(bytes, "settings")
}

func main() {

	// Print the settings message as a JSON string to stdout.
	bytes, err := msettings(settings{"Example Resolver", 1, 10})
	str := string(bytes)
	fmt.Print(str)

	// Read JSON-formatted messages from stdin.
	in := bufio.NewReader(os.Stdin)
	for {
		str, err = in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// End on empty line as input.
		if str == "\n" {
			log.Println("Terminated with empty line for input.")
			return
		}

		b := []byte(str)
		var f interface{}
		err = json.Unmarshal(b, &f)
		m := f.(map[string]interface{})

		msgtype := m["_msgtype"]
		switch msgtype {
		case nil:
			log.Println("Missing _msgtype")
		default:
			log.Println("_msgtype =", msgtype)
		}
	}
}
