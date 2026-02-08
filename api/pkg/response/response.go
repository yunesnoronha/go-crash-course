package response

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse representa uma resposta de erro
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// JSON envia uma resposta JSON
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Error envia uma resposta de erro JSON
func Error(w http.ResponseWriter, statusCode int, message string) {
	errorResponse := ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	}
	JSON(w, statusCode, errorResponse)
}
