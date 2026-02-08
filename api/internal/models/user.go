package models

import (
	"errors"
	"sync"
	"time"
)

// User representa um usuário no sistema
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate valida os dados do usuário
func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("nome é obrigatório")
	}
	if u.Email == "" {
		return errors.New("email é obrigatório")
	}
	if u.Age < 0 || u.Age > 150 {
		return errors.New("idade inválida")
	}
	return nil
}

// UserRepository simula um repositório de usuários em memória
type UserRepository struct {
	users map[string]*User
	mu    sync.RWMutex // Protege acesso concorrente ao map
}

// NewUserRepository cria uma nova instância do repositório
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*User),
	}
}

// FindAll retorna todos os usuários
func (r *UserRepository) FindAll() []*User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users
}

// FindByID busca um usuário por ID
func (r *UserRepository) FindByID(id string) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("usuário não encontrado")
	}
	return user, nil
}

// Create cria um novo usuário
func (r *UserRepository) Create(user *User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID] = user
	return nil
}

// Update atualiza um usuário existente
func (r *UserRepository) Update(id string, user *User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	existingUser, exists := r.users[id]
	if !exists {
		return errors.New("usuário não encontrado")
	}
	if err := user.Validate(); err != nil {
		return err
	}
	user.ID = id
	user.CreatedAt = existingUser.CreatedAt // Preservar created_at original
	r.users[id] = user
	return nil
}

// Delete remove um usuário
func (r *UserRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("usuário não encontrado")
	}
	delete(r.users, id)
	return nil
}
