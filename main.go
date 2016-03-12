package main

import (
	"fmt"
	"net/http"
)

type handler struct {
	wait    chan bool
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
	h.wait <- true
	close(h.wait)
	h.wait = make(chan bool)
}

func (h *handler) polling(w http.ResponseWriter, r *http.Request) {
	<-h.wait
	fmt.Fprintf(w, h.message)
}

func main() {
	ch := make(chan bool)
	h := handler{wait: ch}
	http.HandleFunc("/chat", h.chatHandler)
	http.ListenAndServe(":9000", nil)
}
