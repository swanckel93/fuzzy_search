package handler

import (
	"encoding/json"
	"fmt"
	"github.com/swanckel93/fuzzy_api/models"
	"github.com/swanckel93/fuzzy_api/search"
	"github.com/swanckel93/fuzzy_api/searchCache"
	"github.com/swanckel93/fuzzy_api/storage"
	"github.com/swanckel93/fuzzy_api/utils"
	"io"
	"net/http"
	"time"
	"log"
)

// enableCors sets headers for CORS, including preflight support
func enableCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r)
	if r.Method == http.MethodOptions {
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	text := string(content)
	sentences := utils.SplitIntoSentences(text)

	filename := handler.Filename
	storage.AddFile(filename, sentences)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully"))
}

func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r)
	if r.Method == http.MethodOptions {
		return
	}

	files := storage.ListFiles()
	json.NewEncoder(w).Encode(files)
}

func SearchHandler(w http.ResponseWriter, r *http.Request, cache *searchCache.SearchCache) {
	enableCors(w, r)

	if r.Method == http.MethodOptions {
		return
	}

	// Decode the request body into the SearchRequest model
	var req models.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	fmt.Println("Getting Sentences from storage")

	// Get sentences for the given FileID from storage
	sentences, ok := storage.GetFile(req.FileID)
	if !ok {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	fmt.Println("Got Sentences from storage")

	// Try to get the cached results
	fmt.Println("Attempting to find query in cache")

	cachedResults, found := cache.Get(req.FileID, req.Query)
	if found {
		// Return the cached results if found
		json.NewEncoder(w).Encode(cachedResults)
		fmt.Println("Cache hit! Skipping Fzzy Search...")

		return
	}
	fmt.Println("Not found in Cache. Fuzzysearching...")

	// Perform fuzzy search if not found in cache
	results := search.FuzzySearch(req.Query, sentences)

	// Cache the results for future use
	cache.Set(req.FileID, req.Query, results)
	fmt.Println("Setting search results to cache")

	// Return the results (either from cache or freshly computed)
	json.NewEncoder(w).Encode(results)
}

func ExpandContextHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r)
	if r.Method == http.MethodOptions {
		return
	}

	var req models.ExpandContextRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	sentences, ok := storage.GetFile(req.FileID)
	if !ok || req.Index < 0 || req.Index >= len(sentences) {
		http.Error(w, "Invalid index or file", http.StatusBadRequest)
		return
	}

	context := sentences[req.Index]

	json.NewEncoder(w).Encode(map[string]string{
		"context": context,
	})
}

// Logger middleware for logging requests and response status
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a custom response writer to capture the status code
		writer := &statusCodeWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Call the next handler in the chain
		next.ServeHTTP(writer, r)

		// Log the details of the request
		log.Printf(
			"%s %s %d %s", // log: method, route, status code, duration, user agent
			r.Method,
			r.URL.Path,
			writer.statusCode,
			time.Since(start), // Duration taken for the request
		)
	})
}

// Custom ResponseWriter to capture the status code
type statusCodeWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusCodeWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}