package main

import (
	"WEEK21/internal/entity"
	"WEEK21/internal/repository"
	"WEEK21/internal/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewInMemoryUserRepository()
	useCase := usecase.NewUserUseCase(repo)

	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		var userInput struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		if err := c.BindJSON(&userInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		user := &entity.User{
			Name:  userInput.Name,
			Email: userInput.Email,
		}

		createdUser, err := useCase.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, createdUser)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		rid, _ := strconv.Atoi(id)
		user, err := useCase.GetUserByID(rid)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed: ", err)
	}
	log.Fatal("hello")
}
