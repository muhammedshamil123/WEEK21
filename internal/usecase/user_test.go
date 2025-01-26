package usecase

import (
	"WEEK21/internal/entity"
	"WEEK21/internal/repository/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	useCase := NewUserUseCase(mockRepo)

	t.Run("success", func(t *testing.T) {
		user := &entity.User{Name: "John Doe", Email: "john@example.com"}
		mockRepo.EXPECT().Create(user).Return(user, nil)

		result, err := useCase.CreateUser(user)

		assert.NoError(t, err)
		assert.Equal(t, user, result)
	})

	t.Run("error", func(t *testing.T) {
		user := &entity.User{Name: "John Doe", Email: "john@example.com"}
		mockRepo.EXPECT().Create(user).Return(nil, errors.New("error creating user"))

		result, err := useCase.CreateUser(user)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	useCase := NewUserUseCase(mockRepo)

	t.Run("user found", func(t *testing.T) {
		mockUser := &entity.User{ID: 1, Name: "John Doe", Email: "john@example.com"}
		mockRepo.EXPECT().GetByID(1).Return(mockUser, nil)

		result, err := useCase.GetUserByID(1)

		assert.NoError(t, err)
		assert.Equal(t, mockUser, result)
	})

	t.Run("user not found", func(t *testing.T) {
		mockRepo.EXPECT().GetByID(1).Return(nil, nil)

		result, err := useCase.GetUserByID(1)

		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}
