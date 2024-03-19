package main

import (
	"backend/src/helpers"
	"backend/src/lib"
	"backend/src/middleware"
	"backend/src/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func main() {
	// The database is initialized by init() function in backend/src/lib
	defer lib.CloseDBConnection()
	helpers.LoadEnv()
	r := gin.Default()

	/*
	   Middleware
	*/
	lib.ReplaceLogger(r)
	// middleware.SetCors(r)
	r.Use(middleware.CORSMiddleware())
	goth.UseProviders(middleware.GoogleProvider())
	routes.CreateRouter(r)

	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	store.MaxAge(86400 * 30) // 30 days
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = os.Getenv("ENV") == "PRODUCTION"
	gothic.Store = store

	/*
	   Running service
	*/
	r.Run()
}
