package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/yunesnoronha/go-crash-course/api/internal/models"
	"github.com/yunesnoronha/go-crash-course/api/pkg/response"
)

var userRepo = models.NewUserRepository()

// UsersHandler lida com requisições para /api/v1/users
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listUsers(w, r)
	case http.MethodPost:
		createUser(w, r)
	default:
		response.Error(w, http.StatusMethodNotAllowed, "Método não permitido")
	}
}

// UserHandler lida com requisições para /api/v1/users/:id
func UserHandler(w http.ResponseWriter, r *http.Request) {
	// Extrair ID do path
	id := strings.TrimPrefix(r.URL.Path, "/api/v1/users/")
	if id == "" {
		response.Error(w, http.StatusBadRequest, "ID é obrigatório")
		return
	}

	switch r.Method {
	case http.MethodGet:
		getUser(w, r, id)
	case http.MethodPut:
		updateUser(w, r, id)
	case http.MethodDelete:
		deleteUser(w, r, id)
	default:
		response.Error(w, http.StatusMethodNotAllowed, "Método não permitido")
	}
}

// listUsers lista todos os usuários
func listUsers(w http.ResponseWriter, r *http.Request) {
	users := userRepo.FindAll()
	response.JSON(w, http.StatusOK, users)
}

// getUser busca um usuário por ID
func getUser(w http.ResponseWriter, r *http.Request, id string) {
	user, err := userRepo.FindByID(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, user)
}

// createUser cria um novo usuário
func createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.Error(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	// Gerar ID simples (em produção, use UUID)
	user.ID = generateID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := userRepo.Create(&user); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

// updateUser atualiza um usuário existente
func updateUser(w http.ResponseWriter, r *http.Request, id string) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.Error(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	user.UpdatedAt = time.Now()

	if err := userRepo.Update(id, &user); err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	// Buscar o usuário atualizado para retornar com todos os campos
	updatedUser, err := userRepo.FindByID(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Erro ao buscar usuário atualizado")
		return
	}

	response.JSON(w, http.StatusOK, updatedUser)
}

// deleteUser remove um usuário
func deleteUser(w http.ResponseWriter, r *http.Request, id string) {
	if err := userRepo.Delete(id); err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{
		"message": "Usuário deletado com sucesso",
	})
}

// generateID gera um ID simples (em produção, use UUID)
func generateID() string {
	return time.Now().Format("20060102150405")
}
