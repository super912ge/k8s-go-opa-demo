package app

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) Routes() *gin.Engine {
	router := s.router
	v1 := router.Group("/v1")
	{
		v1.GET("/status", s.ApiStatus())
		user := v1.Group("/user")
		{
			user.POST("", s.CreateUser())
		}
		team := v1.Group("/team")
		{
			team.POST("", s.CreateTeam())
			team.POST("/member", s.CreateTeamMember())
			team.PUT("/manager", s.UpdateManager())
			team.GET("/:id", s.GetTeam())
		}

	}
	return router
}
