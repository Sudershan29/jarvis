package main

import (
  "net/http"
  "backend/src/lib"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  lib.ReplaceLogger(r)

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run()
}