package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "example/xwash-go/util"
  // "fmt"
)

func main() {
  util.InitRedis()
  r := gin.New()
  r.GET("/api/:building", func(c *gin.Context) {
    building := c.Param("building")
    c.JSON(http.StatusOK, gin.H{
      "message": "",
      "data" :util.HGetAll(building),
    })
  })
  r.Run()
}
