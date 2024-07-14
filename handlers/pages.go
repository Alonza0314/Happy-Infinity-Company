package handlers

import (
	"hic/configs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", nil)
}

func GetContact(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.html", nil)
}

func GetSign(c *gin.Context) {
	actionURLSignup, err := configs.GetActionURL("HICserver.addr", "/signup")
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	actionURLSignin, err := configs.GetActionURL("HICserver.addr", "/signin")
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "sign.html", gin.H{
		"ActionURLSignup": actionURLSignup,
		"ActionURLSignin": actionURLSignin,
	})
}
