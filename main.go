package main

import (
	"log"
	"net/http"
	"time"

	"go-fib-service/handlers"
	"go-fib-service/metrics"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/fib", handlers.FibHandler)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.Handle("/metrics", promhttp.Handler())

	finalHandler := metrics.Middleware(mux)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      finalHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Starting server on :8080")
	log.Fatal(server.ListenAndServe())
}
