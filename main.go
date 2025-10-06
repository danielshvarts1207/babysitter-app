package main

import (
	"babysitter-app/handlers"
	"babysitter-app/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db, ctx := repository.Connect("root:my-secret-pw@tcp(127.0.0.1:3306)/babysitterdb?parseTime=true")

	BabysitterRepository := repository.NewBabysitterHandler(db, ctx)
	BabysitterHandler := handlers.NewBabysitterHandler(BabysitterRepository)

	router := gin.Default()

	router.GET("/babysitters", BabysitterHandler.GetAllBabysitters)
	router.GET("/babysitters/:id", BabysitterHandler.GetBabysitterById)
	router.POST("/babysitters", BabysitterHandler.CreateBabysitter)

	router.Run("localhost:8080")
}
