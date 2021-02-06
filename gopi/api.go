package main

import "github.com/gin-gonic/gin"
import "github.com/moonclash/blog-code-examples/gopi/database"

func main() {
	database.CreateDB()
	DatabaseManager := database.New(nil)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/fetch", func(c *gin.Context) {
		short_name, long_name := DatabaseManager.RetrieveDefinition("test")
    c.JSON(200, gin.H{
      "short_name": short_name,
      "long_name": long_name,
    })
	})

	router.POST("/create", func(c *gin.Context) {
		DatabaseManager.InsertDefinition("test", "test definition")
	})
	router.Run()
}