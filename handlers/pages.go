package handlers

import (
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