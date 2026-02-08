package middleware

import (
	"log"
	"net/http"

	"github.com/yunesnoronha/go-crash-course/api/pkg/response"
)

// Recovery middleware para capturar panics
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC: %v", err)
				response.Error(w, http.StatusInternalServerError, "Erro interno do servidor")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
