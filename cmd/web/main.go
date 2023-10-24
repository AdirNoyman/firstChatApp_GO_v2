package main

import (
	"log"
	"net/http"
)

func main() {

	// Start web server
	mux := routes()

	log.Println("Starting web server on port 8080 😎🤘")

	_ = http.ListenAndServe(":8080", mux)

}
