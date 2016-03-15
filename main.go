package main

import (
	"fmt"
	"net/http"
	"sync"
)

type handler struct {
	mu      sync.Mutex
	wait    chan struct{}
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
	// Lock the handler while setting the message and closing the wait
	h.mu.Lock()
	defer h.mu.Unlock()
	body := r.FormValue("body")
	h.message = body
	close(h.wait)
	h.wait = make(chan struct{})

}

func (h *handler) polling(w http.ResponseWriter, r *http.Request) {
	// Lock the handler while accessing the wait chan
	h.mu.Lock()
	wait := h.wait
	h.mu.Unlock()

	<-wait

	// Lock the handler while accessing the message
	h.mu.Lock()
	message := h.message
	h.mu.Unlock()

	fmt.Fprintf(w, message)
}

func newHandler() *handler {
	return &handler{
		wait: make(chan struct{}),
	}
}

func main() {
	h := newHandler()
	http.HandleFunc("/chat", h.chatHandler)
	http.ListenAndServe(":9000", nil)
}
