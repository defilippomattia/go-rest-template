package middleware

import "net/http"

func SetResponseHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set response headers
		w.Header().Set("XXXContent-Type", "application/json") // Example: setting Content-Type

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
