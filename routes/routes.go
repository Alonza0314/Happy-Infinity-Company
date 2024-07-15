package routes

import (
	"hic/handlers"

	"github.com/gin-gonic/gin"
)

func RoutesSetUp(router *gin.Engine) {
	{ // Pages
		router.GET("/", handlers.GetIndex)
		router.GET("/about", handlers.GetAbout)
		router.GET("/contact", handlers.GetContact)
		router.GET("/sign", handlers.GetSign)
		router.GET("/dashboard", handlers.GetDashboard)
	}

	{ // Sign works
		router.POST("/signup", handlers.PostSignup)
		router.POST("/signin", handlers.PostSignin)
	}
}
