package handlers

import (
	"babysitter-app/models/entities"
	"babysitter-app/repository"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type BabysitterHandler struct {
	BabysittersRepository *repository.BabysittersRepository
	validate              *validator.Validate
}

func NewBabysitterHandler(babysittersRepository *repository.BabysittersRepository) *BabysitterHandler {
	return &BabysitterHandler{
		BabysittersRepository: babysittersRepository,
		validate:              validator.New(),
	}
}

func (h *BabysitterHandler) GetAllBabysitters(c *gin.Context) {
	babysitters := h.BabysittersRepository.GetAllBabysitters()

	c.IndentedJSON(http.StatusOK, babysitters)
}

func (h *BabysitterHandler) GetBabysitterById(c *gin.Context) {
	id := c.Param("id")
	babysitter, err := h.BabysittersRepository.GetBabysitterById(c, id)
	switch {
	case err == nil:
		c.IndentedJSON(http.StatusOK, babysitter)
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "Babysitter not found"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}

func (h *BabysitterHandler) CreateBabysitter(c *gin.Context) {
	var babysitter entities.Babysitter
	if err := c.BindJSON(&babysitter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(babysitter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.BabysittersRepository.CreateBabysitter(c, &babysitter)
	c.IndentedJSON(http.StatusCreated, babysitter)
}
