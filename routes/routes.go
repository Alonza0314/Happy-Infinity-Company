package routes

import (
	"hic/handlers"

	"github.com/gin-gonic/gin"
)

func RoutesSetUp(router *gin.Engine) {
	{ // Pages
		router.GET("/", handlers.GetIndex)
	}
}