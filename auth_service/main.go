package main

import (
	"auth-service/config"
	"auth-service/db"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	client := db.ConnectDB()
	defer client.Disconnect(context.TODO())

	router := gin.Default()
	log.Println("Starting auth service on port 3000")
	router.Run(":3000")
}
