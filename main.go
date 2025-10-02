package main

import (
	"babysitter-app/handlers"
	"babysitter-app/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db, ctx := repository.Connect("root:my-secret-pw@tcp(127.0.0.1:3306)/babysitterdb")

	BabysitterHandler := handlers.NewBabysitterHandler(db, ctx)

	router := gin.Default()
	router.GET("/babysitters", BabysitterHandler.GetAllBabysitters)
	router.POST("/babysitters", BabysitterHandler.CreateBabysitter)

	router.Run("localhost:8080")
}
