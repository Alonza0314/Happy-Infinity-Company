package handlers

import (
	"hic/models"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func PostSignup(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	client := models.NewClient(username, email, password)

	var redirectURL string
	if err := models.ProcessSignup(client); err != nil {
		redirectURL = "/sign/?signup=" + url.QueryEscape(err.Error())
	} else {
		redirectURL = "/sign/?signup=true"
	}
	c.Redirect(http.StatusFound, redirectURL)
}
