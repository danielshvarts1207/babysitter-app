package handlers

import (
	"babysitter-app/models/entities"
	"babysitter-app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type BabysitterHandler struct {
	BabysittersRepository repository.BabysittersRepository
	validate              *validator.Validate
}

func NewBabysitterHandler(babysittersRepository repository.BabysittersRepository) *BabysitterHandler {
	return &BabysitterHandler{
		BabysittersRepository: babysittersRepository,
		validate:              validator.New(),
	}
}

func (h *BabysitterHandler) GetAllBabysitters(c *gin.Context) {
	babysitters, _ := gorm.G[entities.Babysitter](h.DB).Find(h.Ctx)

	c.IndentedJSON(http.StatusOK, babysitters)
}

func (h *BabysitterHandler) CreateBabysitter(c *gin.Context) {
	var newBabysitter entities.Babysitter
	if err := c.BindJSON(&newBabysitter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(newBabysitter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.WithContext(h.Ctx).Create(&newBabysitter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newBabysitter)
}
