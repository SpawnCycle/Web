package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Server) MountRoutes() *Server {
	r := gin.Default()
	r.Static("/app", "./build/client")

	api := r.Group("/api")
	{
		api.POST("/api/signup", s.authnController.SignUp)
		api.POST("/api/login", s.authnController.Login)
		api.POST("/api/logout", s.authnController.Logout)
	}

	r.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api") {
			http.ServeFile(c.Writer, c.Request, "./build/client")
		}
		c.Status(http.StatusNotFound)
	})
	s.srv.Handler = r
	return s
}
