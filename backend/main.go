package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Serve static files (frontend)
	staticDir := "../frontend"
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fs)

	// API endpoints
	http.HandleFunc("/api/invoice/lightning", handleLightningInvoice)
	http.HandleFunc("/api/invoice/monero", handleMoneroInvoice)
	http.HandleFunc("/api/status/", handleStatus)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	absPath, err := filepath.Abs(staticDir)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}
	log.Printf("Serving static files from %s", absPath)
	log.Printf("Server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
