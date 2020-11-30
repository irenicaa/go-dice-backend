package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		log.Print("received a request at " + request.URL.String())

		var message string
		if name := request.FormValue("name"); name != "" {
			message = "Hello, user " + name + "!"
		} else {
			message = "Hello, user!"
		}

		writer.Write([]byte(message))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
