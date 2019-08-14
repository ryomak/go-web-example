package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Timeline = func(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "top",
	})
}
