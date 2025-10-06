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

func (r *BabysittersRepository) GetAllBabysitters() []entities.Babysitter {
	babysitters, _ := gorm.G[entities.Babysitter](r.DB).Find(r.Ctx)
	return babysitters
}

func (r *BabysittersRepository) CreateBabysitter(c *gin.Context, babysitter *entities.Babysitter) {
	if err := gorm.G[entities.Babysitter](r.DB).Create(r.Ctx, babysitter); err != nil {
		panic(err)
	}
}

func (r *BabysittersRepository) GetBabysitterById(c *gin.Context, id string) entities.Babysitter {
	babysitter, err := gorm.G[entities.Babysitter](r.DB).Where("ID = ?", id).First(r.Ctx)
	if err != nil {
		panic(err)
	}
	return babysitter
}
