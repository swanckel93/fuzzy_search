package main

import (
	"log"
	"net/http"

	"github.com/swanckel93/fuzzy_api/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("/upload", handler.UploadHandler)
	mux.HandleFunc("/files", handler.ListFilesHandler)
	mux.HandleFunc("/search", handler.SearchHandler)
	mux.HandleFunc("/expand-context", handler.ExpandContextHandler)

	loggedMux := handler.Logger(mux)


	log.Println("Server listening on http://localhost:8080")

	err := http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
