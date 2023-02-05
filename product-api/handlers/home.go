package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Home struct {
	l *log.Logger
}

func NewHome(l *log.Logger) *Home {
	return &Home{l}
}

func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World!")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "An Error Occurred", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%s\n", data)
}
