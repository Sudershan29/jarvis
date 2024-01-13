package main

import (
  "backend/src/lib"
  "backend/src/middleware"
  "backend/src/routes"
  "github.com/gin-gonic/gin"
)

func main() {
  // The database is initialized by init() function in backend/src/lib
  defer lib.CloseDBConnection()
  r := gin.Default()
  lib.ReplaceLogger(r)
  // middleware.SetCors(r)
  r.Use(middleware.CORSMiddleware())
  routes.CreateRouter(r)
  r.Run()
}