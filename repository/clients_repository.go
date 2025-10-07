package repository

import (
	"babysitter-app/models/entities"
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClientsRepository struct {
	DB  *gorm.DB
	Ctx context.Context
}

func NewClientsRepository(db *gorm.DB, ctx context.Context) *ClientsRepository {
	return &ClientsRepository{
		DB:  db,
		Ctx: ctx,
	}
}

func (r *ClientsRepository) GetAllClients() []entities.Client {
	clients, _ := gorm.G[entities.Client](r.DB).Find(r.Ctx)
	return clients
}

func (r *ClientsRepository) CreateClient(c *gin.Context, client *entities.Client) {
	if err := gorm.G[entities.Client](r.DB).Create(r.Ctx, client); err != nil {
		panic(err)
	}
}

func (r *ClientsRepository) GetClientById(c *gin.Context, id string) (entities.Client, error) {
	return gorm.G[entities.Client](r.DB).Where("ID = ?", id).First(r.Ctx)
}
