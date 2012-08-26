package main

import (
	"log"
	"net/http"
	"os"
)

const resp = `<html>
<head>
<title>Example HTTP Server</title>
</head>
<body>
Hello World!
</body>
</html>`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(resp))
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
