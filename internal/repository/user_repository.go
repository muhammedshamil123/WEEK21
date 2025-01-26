package repository

import "WEEK21/internal/entity"

type InMemoryUserRepository struct {
	users map[int]*entity.User
	id    int
}

func NewMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make(map[int]*entity.User)}
}

func (r *InMemoryUserRepository) Create(user *entity.User) (*entity.User, error) {
	r.id++
	user.ID = r.id
	r.users[r.id] = user
	return user, nil
}

func (r *InMemoryUserRepository) GetByID(id int) (*entity.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, nil
	}
	return user, nil
}
