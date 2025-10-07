package main

import (
	"babysitter-app/handlers"
	"babysitter-app/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	db, ctx := repository.Connect("root:my-secret-pw@tcp(127.0.0.1:3306)/babysitterdb?parseTime=true")
	validate := validator.New(validator.WithRequiredStructEnabled())

	UserRepository := repository.NewUserRepository(db, ctx)

	UsersHandler := handlers.NewUsersHandler(UserRepository, validate)

	router := gin.Default()

	router.GET("/users", UsersHandler.GetAllUsers)
	router.GET("/users/:id", UsersHandler.GetUserById)
	router.POST("/users", UsersHandler.CreateUser)

	router.Run("localhost:8080")
}
