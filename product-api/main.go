package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/goodbye", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Goodbye and see you again soon")
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Hello World!")
		data, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "An Error Occurred", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(writer, "%s\n", data)
	})

	// listen from all IPs and serve on port 9091
	http.ListenAndServe(":9091", nil)
}
