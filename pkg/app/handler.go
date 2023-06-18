package app

import (
	"k8s-go-opa/pkg/api"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func (s *Server) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var req api.NewUserRequest

		err := c.ShouldBindJSON(&req)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.userService.CreateUser(req)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "user created",
		})
	}
}

func (s *Server) CreateTeam() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var req api.NewTeamRequest

		err := c.ShouldBindJSON(&req)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.teamService.CreateTeam(req)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "team created",
			"team":   req.Name,
		})
	}
}

func (s *Server) CreateTeamMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var req api.NewTeamMemberRequest

		err := c.ShouldBindJSON(&req)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.teamService.CreateTeamMember(req)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "team member created",
			"team":   req.TeamID,
			"member": req.UserID,
		})
	}
}

func (s *Server) UpdateManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var req api.UpdateManagerRequest

		err := c.ShouldBindJSON(&req)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.teamService.UpdateManager(req)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "manager updated",
			"team":   req.TeamID,
			"member": req.UserID,
		})
	}
}

func (s *Server) GetTeam() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req api.GetTeamRequest

		err := c.ShouldBindUri(&req)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		team, err := s.teamService.GetTeam(req)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "team retrieved",
			"team":   team,
		})
	}
}
