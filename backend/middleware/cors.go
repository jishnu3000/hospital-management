package middleware

import (
    "github.com/gorilla/handlers"
    "net/http"
)

// CORSMiddleware sets up CORS for the application
func CORSMiddleware(next http.Handler) http.Handler {
    return handlers.CORS(
        handlers.AllowedOrigins([]string{"http://127.0.0.1:5500"}), // Allow frontend requests
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Allow HTTP methods
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Allow necessary headers
    )(next)
}
