package handlers

import (
	"encoding/json"
	"hic/configs"
	"hic/models"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-contrib/sessions"
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

func PostSignin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	client := models.NewClient(username, "", password)

	if err := models.ProcessSignin(client); err != nil {
		c.Redirect(http.StatusFound, "/sign/?signin="+url.QueryEscape(err.Error()))
		c.Abort()
		return
	}

	// set session
	userInfo := models.UserInfo{Username: username, SigninTime: time.Now()}
	userJSON, err := json.Marshal(userInfo)
	if err != nil {
		log.Println("server error => postsignin function error\n\t" + err.Error())
		c.Redirect(http.StatusFound, "/sign/?signin="+url.QueryEscape("server error.\n服務器錯誤"))
		c.Abort()
		return
	}
	session := sessions.Default(c)
	userid := models.GenerateHash(userInfo)
	session.Set(userid, string(userJSON))
	session.Save()

	// set cookie "userid"
	cookieTimeout, err := configs.GetSessionCookieTimeout("cookie.timeout")
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusFound, "/sign/?signin="+url.QueryEscape("server error.\n服務器錯誤"))
		c.Abort()
		return
	}
	c.SetCookie("userid", userid, cookieTimeout*60, "/", "", false, true)
	/*
		1. name：cookie 的名称。
		2. value：cookie 的值。
		3. maxAge：cookie 的过期时间（以秒为单位），如果为正数，则表示 cookie 在指定的秒数后过期；如果为负数，则表示 cookie 在浏览器关闭后过期；如果为 0，则表示立即删除该 cookie。在示例中，3600 表示 cookie 在 3600 秒（即 1 小时）后过期。
		4. path：cookie 的作用路径。指定 cookie 生效的路径，浏览器只会向该路径发送 cookie。在示例中，"/" 表示整个网站都可以接收到该 cookie。
		5. domain：cookie 的作用域。指定哪些域名可以接收到 cookie。如果为空字符串，则表示只有设置该 cookie 的域名可以接收到；如果为 nil，则表示使用当前请求的域名。在示例中，"" 表示使用当前请求的域名。
		6. secure：是否只在 HTTPS 连接下传输 cookie。如果为 true，则表示只在 HTTPS 连接下传输；如果为 false，则表示在 HTTP 和 HTTPS 连接下都传输。在示例中，false 表示在 HTTP 和 HTTPS 连接下都传输。
		7. httpOnly：是否限制 cookie 只能通过 HTTP 或 HTTPS 协议传输，而不能通过 JavaScript 访问。如果为 true，则表示只能通过 HTTP 或 HTTPS 协议传输；如果为 false，则表示可以通过 JavaScript 访问。在示例中，true 表示只能通过 HTTP 或 HTTPS 协议传输。
	*/

	c.Redirect(http.StatusFound, "/dashboard")
}

func PostPwfind(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	client := models.NewClient(username, email, "")

	if err := models.ProcessPwfind(client); err != nil {
		c.Redirect(http.StatusFound, "/sign/pwfind?pwfind=" + err.Error())
		c.Abort()
		return
	}

	resetInfo := models.ResetInfo{Username: username, Email: email}
	resetJSON, err := json.Marshal(resetInfo)
	if err != nil {
		log.Println("server error => postpwfind function error\n\t" + err.Error())
		c.Redirect(http.StatusFound, "/sign/pwfind?pwfind=" + url.QueryEscape("server error find password fail.\n\n服務器錯誤密碼找回失敗"))
		c.Abort()
		return
	}
	session := sessions.Default(c)
	session.Set("resetid", string(resetJSON))
	session.Save()

	c.Redirect(http.StatusFound, "/sign/pwreset")
}
