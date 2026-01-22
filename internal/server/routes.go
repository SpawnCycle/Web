package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) MountRoutes() *Server {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"data": "ok"}) })

	s.srv.Handler = r
	return s
}
