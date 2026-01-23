package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) MountRoutes() *Server {
	r := gin.Default()
	r.GET("/users", s.userController.ReadAllUsers)
	r.GET("/users/:id", s.userController.ReadUserByID)

	s.srv.Handler = r
	return s
}
