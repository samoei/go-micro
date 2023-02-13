package main

import (
	"context"
	"github.com/gorilla/mux"
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

	//create a products endpoint
	products := handlers.NewProducts(logger)

	// Create a server mux AKA router
	router := mux.NewRouter()

	// Get Route
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", products.GetProducts)

	// PUT Route
	putRoute := router.Methods(http.MethodPut).Subrouter()
	putRoute.HandleFunc("/{id:[0-9]+}", products.UpdateProducts)

	//register handlers
	//router.Handle("/", products)

	// server properties
	server := http.Server{
		Addr:         ":9091",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//start the server on a different thread
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// make a channel
	sigChan := make(chan os.Signal)

	//listen for interupt or kill signal on the channel
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// register a listener on the main thread
	sig := <-sigChan

	logger.Println("Received a terminate request of type:", sig, "\nGraceful shutdown...")

	//allow for 30 for graceful shutdown
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
