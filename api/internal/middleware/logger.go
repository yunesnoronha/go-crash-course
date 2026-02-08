package middleware

import (
	"log"
	"net/http"
	"time"
)

// responseWriter wrapper para capturar o status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Logger middleware para logging de requisições
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrapper para capturar status code
		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		// Processar requisição
		next.ServeHTTP(wrapped, r)

		// Log da requisição
		duration := time.Since(start)
		log.Printf(
			"[%s] %s %s - Status: %d - Duration: %v",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			wrapped.statusCode,
			duration,
		)
	})
}
