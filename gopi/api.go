package main

import (
  "github.com/gin-gonic/gin"
  "github.com/moonclash/blog-code-examples/gopi/database"
  "fmt"
)

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
    fmt.Println(c.Params)
    DatabaseManager.InsertDefinition("test", "test definition")
  })
  router.Run()
}