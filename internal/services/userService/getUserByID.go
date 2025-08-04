package userService

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// GetUserByID получает пользователя по external_id
func (s *UserService) GetUserByID(c *gin.Context) {
	userIDStr := c.Param("id")

	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id format"})
		return
	}

	// Получаем всех пользователей и ищем по ExternalID
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get users"})
		return
	}

	for _, user := range *users {
		if user.ExternalID == userID {
			c.JSON(http.StatusOK, gin.H{
				"external_id": user.ExternalID.String(),
				"created_at":  user.CreatedAt,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
} 