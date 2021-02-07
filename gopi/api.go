package main

import (
  "encoding/json"
  "github.com/gin-gonic/gin"
  "github.com/moonclash/blog-code-examples/gopi/database"
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

  router.GET("/fetch/:short", func(c *gin.Context) {
    short_def := c.Param("short")
    short_name, long_name := DatabaseManager.RetrieveDefinition(short_def)
    c.JSON(200, gin.H{
      "short_name": short_name,
      "long_name": long_name,
    })
  })

  router.POST("/create", func(c *gin.Context) {
    decoder := json.NewDecoder(c.Request.Body)
    var definitionStruct struct {
      Short string
      Long string
    }
    decoder.Decode(&definitionStruct)
    DatabaseManager.InsertDefinition(definitionStruct.Short, definitionStruct.Long)
  })
  router.Run()
}