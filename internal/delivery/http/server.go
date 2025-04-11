package http

import (
	"log"
	"net/http"
)

func StartServer(router *Router, addr string) {
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
