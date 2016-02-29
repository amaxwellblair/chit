package main

import (
	"fmt"
	"net/http"
)

var x bool
var body string

func rootHandler(w http.ResponseWriter, r *http.Request) {
	body = r.FormValue("body")
	x = true
	fmt.Println("yea we got it")
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	x = false
	for x == false {

	}
	fmt.Fprintf(w, body)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/chat/", chatHandler)
	http.ListenAndServe(":9000", nil)
}
