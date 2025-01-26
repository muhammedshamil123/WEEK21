package usecase

import (
	"WEEK21/internal/entity"
	"WEEK21/internal/repository"
)

type UserUseCase interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByID(id int) (*entity.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (u *userUseCase) CreateUser(user *entity.User) (*entity.User, error) {
	return u.repo.Create(user)
}

func (u *userUseCase) GetUserByID(id int) (*entity.User, error) {
	return u.repo.GetByID(id)
}
