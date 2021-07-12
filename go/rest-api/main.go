package main

import (
	"log"
	"net/http"

	"rest-api/routes"
)

func main() {
	// This is basically binding a Gin server on a normal HTTP server
	router := routes.Setup()
	server := &http.Server{
		Addr:    "localhost:5000",
		Handler: router,
	}
	log.Printf("[info] http server is listening on http://localhost:5000")
	server.ListenAndServe()
}
