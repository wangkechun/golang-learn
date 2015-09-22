package web

import (
	"log"
	"net/http"
)

// Handler is the main handler
type Handler struct {
}

// New a Handler
func New() *Handler {
	return &Handler{}
}

// ServerHTTP ...
func (p *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	w.Write([]byte("aaa"))
}

// Run server
func (p *Handler) Run(addr string) error {
	return http.ListenAndServe(addr, p)
}
