package main

import (
	"context"
	"hic/configs"
	"hic/handlers"
	"hic/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	router.Use(handlers.PwresetRedirect)

	routes.RoutesSetUp(router)

	// router.Run(":8080")
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Println("@ Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("server error => main function error\n\t", err.Error())
		}
	}()

	<-stop
	log.Println("@ Shutting down server...")
	
	gracefulTimeout, err := configs.GetGracefulTimeout()
	if err != nil {
		log.Println(err)
		gracefulTimeout = 10
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(gracefulTimeout) * time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("server error => main function error\n\t", err.Error())
	}

	log.Println("@ Server gracefully stopped")
}
