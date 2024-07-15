package main

import (
	"hic/configs"
	"hic/handlers"
	"hic/routes"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

/*
⣴⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣴⡷
⠈⣿⣷⣦⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣶⣿⣿
⠀⢸⣿⣿⣿⣿⣷⣆⣀⣀⣀⣀⣀⣾⣿⣿⣿⣿⡇
⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇
⠀⠀⠿⢿⣿⣿⣿⣿⡏⡀⠀⡙⣿⣿⣿⣿⣿⠛
⠀⠀⠀⣿⣿⣿⡿⠟⠷⣅⣀⠵⠟⢿⣿⣿⣿⡆
⠀⠀⠀⣿⣿⠏⢲⣤⠀⠀⠀⠀⢠⣶⠙⣿⣿⠃
⠀⠀⠀⠘⢿⡄⠈⠃⠀⢐⢔⠀⠈⠋⢀⡿⠋
⠀⠀⠀⢀⢀⣼⣷⣶⣤⣤⣭⣤⣴⣶⣍
⠀⠀⠀⠈⠈⣈⢰⠿⠛⠉⠉⢻⢇⠆⣁⠁
⠀⠀⠀⠀⠀⠑⢸⠉⠀⠀⠀⠀⠁⡄⢘⣽⣿
⠀⠀⠀⠀⠀⠀⡜⠀⠀⢰⡆⠀⠀⠻⠛⠋
⠀⠀⠀⠀⠀⠀⠑⠒⠒⠈⠈⠒⠒⠊
||||||||||||||||||||||||||||||
||||||||||||KUROMI||||||||||||
|||||BLESSING|||||PROGRAM|||||
||||||||||||||||||||||||||||||
*/

func main() {
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	store := cookie.NewStore([]byte("dufeng0314"))
	router.Use(sessions.Sessions("mysession", store))

	sessionTimeout, err := configs.GetSessionCookieTimeout("session.timeout")
	if err != nil {
		log.Println(err)
		return
	}
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: int(time.Duration(sessionTimeout) * time.Minute),
	})

	router.Use(handlers.SigninRedirect)
	router.Use(handlers.NoneSigninRedirect)

	routes.RoutesSetUp(router)

	router.Run(":8080")
}
