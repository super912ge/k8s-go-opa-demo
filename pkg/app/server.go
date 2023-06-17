package app

import (
	"k8s-go-opa/pkg/api"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router      *gin.Engine
	userService api.UserService
	teamService api.TeamService
}

func NewServer(router *gin.Engine, userService api.UserService, teamService api.TeamService) *Server {
	return &Server{
		router:      router,
		userService: userService,
		teamService: teamService,
	}
}

func (s *Server) Run() error {
	r := s.Routes()
	err := r.Run(":8081")
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
