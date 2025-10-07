package main

import (
	"babysitter-app/handlers"
	"babysitter-app/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db, ctx := repository.Connect("root:my-secret-pw@tcp(127.0.0.1:3306)/babysitterdb?parseTime=true")

	BabysitterRepository := repository.NewBabysitterHandler(db, ctx)
	ClientRepository := repository.NewClientsRepository(db, ctx)

	BabysitterHandler := handlers.NewBabysitterHandler(BabysitterRepository)
	ClientHandler := handlers.NewClientHandler(ClientRepository)

	router := gin.Default()

	router.GET("/clients", ClientHandler.GetAllClients)
	router.GET("/clients/:id", ClientHandler.GetClientById)
	router.POST("/clients", ClientHandler.CreateClient)

	router.GET("/babysitters", BabysitterHandler.GetAllBabysitters)
	router.GET("/babysitters/:id", BabysitterHandler.GetBabysitterById)
	router.POST("/babysitters", BabysitterHandler.CreateBabysitter)

	router.Run("localhost:8080")
}
