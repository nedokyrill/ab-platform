package userService

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUsers получает список всех пользователей
func (s *UserService) GetUsers(c *gin.Context) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get users"})
		return
	}

	var response []gin.H
	for _, user := range *users {
		response = append(response, gin.H{
			"external_id": user.ExternalID.String(),
			"created_at":  user.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"users": response,
		"count": len(response),
	})
} 