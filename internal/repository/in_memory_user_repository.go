package repository

import (
	"WEEK21/internal/entity"
	"errors"
	"sync"
)

type inMemoryUserRepository struct {
	users map[int]*entity.User
	mu    sync.RWMutex
}

func NewInMemoryUserRepository() UserRepository {
	return &inMemoryUserRepository{
		users: make(map[int]*entity.User),
	}
}

func (r *inMemoryUserRepository) Create(user *entity.User) (*entity.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Simulate auto-incrementing ID
	user.ID = len(r.users) + 1
	r.users[user.ID] = user
	return user, nil
}

func (r *inMemoryUserRepository) GetByID(id int) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
