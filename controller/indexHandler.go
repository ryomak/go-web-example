package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Index(c *gin.Context) {
	info, _ := GetSessionInfo(c)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "top",
		"session": info,
	})
}
