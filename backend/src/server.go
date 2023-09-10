package main

import (
  "fmt"
  "time"
  "net/http"
  "go.uber.org/zap"
  "github.com/gin-gonic/gin"
  ginzap "github.com/gin-contrib/zap"
  helper "backend/src/helpers"
)

func main() {
  r := gin.Default()

  /*
    Referencing https://github.com/gin-contrib/zap

    Building a zap logger

  */
  config := zap.NewProductionConfig()
  config.OutputPaths = []string{fmt.Sprintf("/usr/src/app/logs/%s.log", helper.GetEnv("env", "development")), "stdout"}
  zapLogger, _ := config.Build()
  r.Use(ginzap.GinzapWithConfig(zapLogger, &ginzap.Config{
    TimeFormat: time.RFC3339,
    UTC: true,
  }))

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run()
}