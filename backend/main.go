package main

import (
	"log"
	"net/http"

	"github.com/swanckel93/fuzzy_api/handlers"
	"github.com/swanckel93/fuzzy_api/searchCache"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swanckel93/fuzzy_api/docs" // required for generated docs
)

func main() {
	mux := http.NewServeMux()
	cache := searchCache.NewSearchCache(50)

	// Define routes
	mux.Handle("/docs/", httpSwagger.WrapHandler)
	mux.HandleFunc("/upload", handler.UploadHandler)
	mux.HandleFunc("/files", handler.ListFilesHandler)
	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		handler.SearchHandler(w, r, cache)
	})
	mux.HandleFunc("/expand-context", handler.ExpandContextHandler)

	loggedMux := handler.Logger(mux)


	log.Println("Server listening on http://localhost:8080")
	log.Println("")
	log.Println("SWAGGER DOCS HERE: http://localhost:8080/docs/index.html")

	err := http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
