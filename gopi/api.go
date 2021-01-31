package main

import "github.com/gin-gonic/gin"
import "database"

func main() {
	database.CreateDB()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}