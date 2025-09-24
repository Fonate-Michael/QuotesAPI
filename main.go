package main

import (
	"app/config"
	"app/routes"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Bitch API!")
	config.ConnectToDB()
	r := gin.Default()

	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"POST", "GET", "PATCH", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true

	r.Use(cors.New(config))

	routes.ServerRoutes(r)
	r.Run(":8080")
}
