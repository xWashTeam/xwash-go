package main

import (
  // "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
  "xwash/go/util"
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

  // machine := util.Machine{
  //   "d19_E_1.5",
  //   "d19",
  //   "uClean",
  //   "东十九东边1.5楼",
  //   "https://q.ujing.com.cn/ucqrc/index.html?cd=0000000000000A0007555202104170176799",
  // }
  // qr := util.Check(machine)
  // fmt.Println(qr)
}
