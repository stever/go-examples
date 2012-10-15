package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var s string
	var err error
	for {
		s, err = in.ReadString('\n')

		if err != nil {
			return
		}

		if s == "\n" {
			return
		}

		b := []byte(s)
		var f interface{}
		err = json.Unmarshal(b, &f)
		m := f.(map[string]interface{})

		msgtype := m["_msgtype"]
		if msgtype == nil {
			return
		}

		switch msgtype {
		case "settings":
			// {"_msgtype":"settings","name":"test","weight":1,"targettime":10}
			fmt.Println("Settings:")
			fmt.Println("name =", m["name"])
			fmt.Println("weight =", m["weight"])
			fmt.Println("targettime =", m["targettime"])
		case "rq":
			fmt.Println("rq")
		case "results":
			fmt.Println("results")
		}
	}
}
