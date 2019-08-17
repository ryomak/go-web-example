package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ryomak/go-web-example/controller"
	"github.com/ryomak/go-web-example/util"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*")

	server := &controller.Server{
		DB: util.GormConnect(),
	}

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/", server.Index)
	router.GET("/login", server.Login)
	router.POST("/login", server.LoginPost)
	router.GET("/logout", server.Logout)
	router.GET("/signup", server.SignUp)
	router.POST("/signup", server.SignUpPost)

	privateRouter := router.Group("/me")
	privateRouter.Use(controller.AuthMiddle())
	privateRouter.GET("/mypage", server.Me)

	router.Run(":8080")
}
