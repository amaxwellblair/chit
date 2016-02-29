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
	ch := make(chan bool)
	go streaming(w, ch)
	<-ch
}

func streaming(w http.ResponseWriter, ch chan bool) {
	for x == false {
		// time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(w, body)
	ch <- true
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/chat", chatHandler)
	http.ListenAndServe(":9000", nil)
}
