package http

import (
	"WEEK21/internal/entity"
	"WEEK21/internal/usecase/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
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

	// Mock user and behavior
	mockUser := &entity.User{ID: 1, Name: "John", Email: "john@example.com"}
	mockUseCase.EXPECT().CreateUser(gomock.Any()).Return(mockUser, nil).Times(1)

	// Corrected request body
	requestBody := `{
        "name": "John",
        "email": "john@example.com",
        "password": "securepassword"
    }`

	// Create and send the request
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/users", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	// Assert the status code and response
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `"id":1`)
	assert.Contains(t, w.Body.String(), `"name":"John"`)
}
