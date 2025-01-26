package repository

import "WEEK21/internal/entity"

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	GetByID(id int) (*entity.User, error)
}
