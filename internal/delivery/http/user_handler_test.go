package http

import (
	"WEEK21/internal/entity"
	"WEEK21/internal/usecase/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockUserUseCase(ctrl)
	handler := &UserHandler{UseCase: mockUseCase}

	router := gin.Default()
	router.POST("/users", handler.CreateUser)

	mockUser := &entity.User{ID: 1, Name: "John", Email: "john@example.com"}
	mockUseCase.EXPECT().CreateUser(gomock.Any()).Return(mockUser, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/users", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
