package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type APIRouter struct {
	router *gin.Engine
}

func CreateRouter() *APIRouter {
	router := gin.Default()
}