package main

import (
	"fmt"
	"net/http"
	"time"
)

var x bool
var body string

// rootHandler recieves messages
func rootHandler(w http.ResponseWriter, r *http.Request) {
	body = r.FormValue("body")
	x = true
	fmt.Println("Message received: ", body)
}

// chatHandler sends messages to clients connected to the chat room
func chatHandler(w http.ResponseWriter, r *http.Request) {
	x = false
	for x == false {
		// Sleep will prevent the long polling from blocking
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Fprintf(w, body)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/chat", chatHandler)
	http.ListenAndServe(":9000", nil)
}
