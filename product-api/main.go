package main

import (
	"log"
	"main/handlers"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	home := handlers.NewHome(logger)

	serverMux := http.NewServeMux()
	serverMux.Handle("/", home)

	// listen from all IPs and serve on port 9091
	http.ListenAndServe(":9091", serverMux)
}
