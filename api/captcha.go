package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

func GetCaptcha(c *gin.Context) {
	id, imgData, code, err := MakeCaptcha()
	if err != nil {
		log.Println("server error => getcaptcha function error\n\t" + err.Error())
		c.String(http.StatusInternalServerError, "server error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"code":  code,
		"image": imgData,
	})
}

func MakeCaptcha() (string, string, string, error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	/*
		width：寬度。
		height：高度。
		noiseCount：干擾線數量。
		showLineOptions：干擾線的可見度，0是不顯示，1是顯示。
		length：驗證碼長度。
	*/

	cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)

	id, b64s, code, err := cp.Generate()
	if err != nil {
		return "", "", "", err
	}

	return id, b64s, code, nil
}
