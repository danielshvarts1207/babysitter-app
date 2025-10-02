package repository

import (
	"babysitter-app/models/entities"
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BabysittersRepository struct {
	DB  *gorm.DB
	Ctx context.Context
}

func NewBabysitterHandler(db *gorm.DB, ctx context.Context) *BabysittersRepository {
	return &BabysittersRepository{
		DB:  db,
		Ctx: ctx,
	}
}

func (r *BabysittersRepository) GetAllBabysitters(c *gin.Context) []entities.Babysitter {
	babysitters, _ := gorm.G[entities.Babysitter](r.DB).Find(r.Ctx)
	return babysitters
}

func (r *BabysittersRepository) CreateBabysitter(c *gin.Context, newBabysitter entities.Babysitter) {
	gorm.G[entities.Babysitter](r.DB).Create(r.Ctx, &newBabysitter)
}
