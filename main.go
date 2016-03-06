package main

import (
	"fmt"
	"net/http"
)

type handler struct {
	message chan string
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
	h.message <- r.FormValue("body")
	close(h.message)
	h.message = make(chan string)
}

func (h *handler) polling(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, <-h.message)
}

func main() {
	ch := make(chan string)
	h := handler{message: ch}
	http.HandleFunc("/chat", h.chatHandler)
	http.ListenAndServe(":9000", nil)
}
