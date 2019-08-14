package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ryomak/go-web-example/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*")
	router.GET("/", controller.Index)
	router.Run(":8080")
}
