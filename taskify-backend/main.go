package main

import (
	"log"
	"net/http"
	"os"
	"taskify/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Taskify index endpoint"})
	})

	routes.TodoRoutes(router)

	if err := router.Run(":" + os.Getenv("SERVER_PORT")); err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
