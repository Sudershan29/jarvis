package lib

import (
  "fmt"
  "time"
  "go.uber.org/zap"
  "github.com/gin-gonic/gin"
  ginzap "github.com/gin-contrib/zap"
  helper "backend/src/helpers"
)

func ReplaceLogger(r *gin.Engine){
  /*
    Referencing https://github.com/gin-contrib/zap

    Building a zap logger

  */
  config := zap.NewProductionConfig()
  config.OutputPaths = []string{fmt.Sprintf("/usr/src/app/logs/%s.log", helper.GetEnvWithDefault("env", "development")), "stdout"}
  zapLogger, _ := config.Build()
  r.Use(ginzap.GinzapWithConfig(zapLogger, &ginzap.Config{
    TimeFormat: time.RFC3339,
    UTC: true,
  }))
}