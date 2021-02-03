package main

import "github.com/gin-gonic/gin"
import "github.com/moonclash/blog-code-examples/gopi/database"

func main() {
	database.CreateDB()
	DatabaseManager := database.DBManager.new("foo")
	DatabaseManager.initialize()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}