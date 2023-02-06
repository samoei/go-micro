package main

import (
	"context"
	"log"
	"main/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//handlers do need a logger
	logger := log.New(os.Stdout, "product-api:", log.LstdFlags)

	//create a home page handler
	products := handlers.NewProducts(logger)

	// Create a server mux
	handler := http.NewServeMux()

	//register handlers
	handler.Handle("/", products)

	server := http.Server{
		Addr:         ":9091",
		Handler:      handler,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	logger.Println("Received a terminate request of type:", sig, "\nGraceful shutdown...")

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
