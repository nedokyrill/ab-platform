package userService

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nedokyrill/ab-platform/internal/models"
	"net/http"
	"time"
)

// CreateUser создает нового пользователя
func (s *UserService) CreateUser(c *gin.Context) {
	var request struct {
		ExternalID string `json:"external_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	// Парсим external_id
	externalID, err := uuid.Parse(request.ExternalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid external_id format"})
		return
	}

	// Создаем нового пользователя
	user := models.UserModel{
		ExternalID: externalID,
		CreatedAt:  time.Now(),
	}

	err = s.userRepo.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"external_id": user.ExternalID.String(),
		"created_at":  user.CreatedAt,
	})
} 