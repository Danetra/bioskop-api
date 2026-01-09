package main

import (
	"bioskop-api/config"
	"bioskop-api/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.DbConfig()

	routes.BioskopRoutes(router)
	fmt.Println("Server running at http://localhost:8080")
	router.Run(":8080")
}
