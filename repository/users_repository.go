package repository

import (
	dtos "babysitter-app/models/dtos/babysitters"
	"babysitter-app/models/entities"
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UsersRepository struct {
	DB  *gorm.DB
	Ctx context.Context
}

func NewUserRepository(db *gorm.DB, ctx context.Context) *UsersRepository {
	return &UsersRepository{
		DB:  db,
		Ctx: ctx,
	}
}

func (r *UsersRepository) GetAllUsers() []entities.User {
	users, _ := gorm.G[entities.User](r.DB).Find(r.Ctx)
	return users
}

func (r *UsersRepository) CreateUser(c *gin.Context, user *dtos.CreateUserDto) entities.User {
	userEntity := entities.User{
		Name:         user.Name,
		FamilyName:   user.FamilyName,
		Email:        user.Email,
		Role:         entities.Role(user.Role),
		PasswordHash: user.PasswordHash,
	}
	if err := gorm.G[entities.User](r.DB).Create(r.Ctx, &userEntity); err != nil {
		panic(err)
	}
	return userEntity
}

func (r *UsersRepository) GetUserById(c *gin.Context, id string) (entities.User, error) {
	return gorm.G[entities.User](r.DB).Where("ID = ?", id).First(r.Ctx)
}
