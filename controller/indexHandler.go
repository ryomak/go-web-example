package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Index(c *gin.Context) {
	i, _ := c.Get("authUser")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "top",
		"session": i,
	})
}
