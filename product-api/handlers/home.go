package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Home struct
type Home struct {
	l *log.Logger
}

// NewHome initialise new home struct with a logger
func NewHome(l *log.Logger) *Home {
	return &Home{l}
}

// implement the ServerHTTP interface
func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Welcome to the Home Page")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "An Error Occurred", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%s\n", data)
}
