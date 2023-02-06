package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type About struct {
	l *log.Logger
}

func NewAbout(l *log.Logger) *About {
	return &About{l}
}

func (a *About) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("Welcome to the About Page")

	data, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "An Error Occurred", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "%s", data)

}
