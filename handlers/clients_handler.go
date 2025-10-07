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

type ClientHandler struct {
	ClientsRepository *repository.ClientsRepository
	validate          *validator.Validate
}

func NewClientHandler(clientsRepository *repository.ClientsRepository) *ClientHandler {
	return &ClientHandler{
		ClientsRepository: clientsRepository,
		validate:          validator.New(),
	}
}

func (h *ClientHandler) GetAllClients(c *gin.Context) {
	clients := h.ClientsRepository.GetAllClients()

	c.IndentedJSON(http.StatusOK, clients)
}

func (h *ClientHandler) GetClientById(c *gin.Context) {
	id := c.Param("id")
	client, err := h.ClientsRepository.GetClientById(c, id)
	switch {
	case err == nil:
		c.IndentedJSON(http.StatusOK, client)
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}

func (h *ClientHandler) CreateClient(c *gin.Context) {
	var client entities.Client
	if err := c.BindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.ClientsRepository.CreateClient(c, &client)
	c.IndentedJSON(http.StatusCreated, client)
}
