package handlers

import (
	dtos "babysitter-app/models/dtos/babysitters"
	"babysitter-app/repository"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UsersHandler struct {
	UserRepository *repository.UsersRepository
	validate       *validator.Validate
}

func NewUsersHandler(babysittersRepository *repository.UsersRepository, validate *validator.Validate) *UsersHandler {
	return &UsersHandler{
		UserRepository: babysittersRepository,
		validate:       validate,
	}
}

func (h *UsersHandler) GetAllUsers(c *gin.Context) {
	users := h.UserRepository.GetAllUsers()

	c.IndentedJSON(http.StatusOK, users)
}

func (h *UsersHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := h.UserRepository.GetUserById(c, id)
	switch {
	case err == nil:
		c.IndentedJSON(http.StatusOK, user)
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}

func (h *UsersHandler) CreateUser(c *gin.Context) {
	var user dtos.CreateUserDto
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userEntity := h.UserRepository.CreateUser(c, &user)
	c.IndentedJSON(http.StatusCreated, userEntity)
}
