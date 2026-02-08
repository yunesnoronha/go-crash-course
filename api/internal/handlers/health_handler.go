package handlers

import (
	"net/http"
	"time"

	"github.com/yunesnoronha/go-crash-course/api/pkg/response"
)

// HealthCheck retorna o status de saúde da aplicação
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Error(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	health := map[string]interface{}{
		"status":    "ok",
		"timestamp": time.Now().Format(time.RFC3339),
		"service":   "go-crash-course-api",
	}

	response.JSON(w, http.StatusOK, health)
}
