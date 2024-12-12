package main

import (
	"log"
	"net/http"

	"github.com/ammyy9908/codejudge/internal/submission/handlers"
)

// CORS middleware function
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")              // Allow all origins (change as needed for security)
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS") // Allow specific HTTP methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")  // Allow specific headers

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass to the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Wrap your handler with the CORS middleware
	http.Handle("/submit", corsMiddleware(http.HandlerFunc(handlers.SubmitHandler)))

	log.Println("Starting submission-service on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
