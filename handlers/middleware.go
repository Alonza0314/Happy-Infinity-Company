package handlers

import (
	"encoding/json"
	"hic/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SigninRedirect(c *gin.Context) {
	if c.Request.Method == "GET" && IsSignin(c) && c.Request.URL.Path == "/sign/" {
		c.Redirect(http.StatusFound, "/dashboard/")
		c.Abort()
		return
	}
	c.Next()
}

func NoneSigninRedirect(c *gin.Context) {
	if c.Request.Method == "GET" && !IsSignin(c) && !CheckPathDoNotNeedSignin(c) {
		c.Redirect(http.StatusFound, "/sign/")
		c.Abort()
		return
	}
	c.Next()
}

func PwresetRedirect(c *gin.Context) {
	if c.Request.URL.Path == "/sign/pwreset" {
        session := sessions.Default(c)
        resetid := session.Get("resetid")

        if resetid == nil {
            c.Redirect(http.StatusFound, "/sign/pwfind")
            c.Abort()
            return
        }
    }

    c.Next()
}

func CheckPathDoNotNeedSignin(c *gin.Context) bool {
	paths := []string{
		"/",
		"/about",
		"/contact",
		"/sign/",
		"/sign/pwfind",
		"/sign/pwreset",	

		"/signup",
		"/signin",
		"/pwfind",
		"/pwreset",

		"/api/captcha",
	}
	for _, path := range paths {
		if c.Request.URL.Path == path {
			return true
		}
	}
	return false
}

func IsSignin(c *gin.Context) bool {
	cookie, err := c.Request.Cookie("userid")
	if err != nil {
		return false
	}

	session := sessions.Default(c)
	userid := cookie.Value
	userJSON := session.Get(userid)
	if userJSON == nil {
		return false
	}

	var userInfo models.UserInfo
	json.Unmarshal([]byte(userJSON.(string)), &userInfo)

	return time.Since(userInfo.SigninTime) <= 5*time.Minute
}
