package routes

import (
	"hic/api"
	"hic/handlers"

	"github.com/gin-gonic/gin"
)

func RoutesSetUp(router *gin.Engine) {
	{ // Pages
		router.GET("/", handlers.GetIndex)
		router.GET("/about", handlers.GetAbout)
		router.GET("/contact", handlers.GetContact)
		signGroup := router.Group("/sign")
		{
			signGroup.GET("/", handlers.GetSign)
			signGroup.GET("/pwfind", handlers.GetPwfind)
			// signGroup.GET("/resetpw", )
		}
		dashboardGroup := router.Group("/dashboard")
		{
			dashboardGroup.GET("/", handlers.GetDashboard)
		}
	}

	{ // Sign works
		router.POST("/signup", handlers.PostSignup)
		router.POST("/signin", handlers.PostSignin)
	}

	{
		// api
		apiGroup := router.Group("/api")
		{
			apiGroup.GET("/captcha", api.GetCaptcha)
		}
	}
}
