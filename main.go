package main

import (
	"fmt"
	"net/http"
)

type handler struct {
	ch      chan string
	message string
}

func (h *handler) chatHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.polling(w, r)
	case "POST":
		h.broadcast(w, r)
	}
}

func (h *handler) broadcast(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	h.message = body
	h.ch <- body
	close(h.ch)
	h.ch = make(chan string)
}

func (h *handler) polling(w http.ResponseWriter, r *http.Request) {
	<-h.ch
	fmt.Fprintf(w, h.message)
}

func main() {
	ch := make(chan string)
	h := handler{ch: ch}
	http.HandleFunc("/chat", h.chatHandler)
	http.ListenAndServe(":9000", nil)
}
