package assignmentService

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// GetAssignedVariant получает уже назначенный пользователю вариант
func (s *AssignmentService) GetAssignedVariant(c *gin.Context) {
	userIDStr := c.Query("user_id")
	experimentIDStr := c.Query("experiment_id")

	if userIDStr == "" || experimentIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id and experiment_id are required"})
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id format"})
		return
	}

	experimentID, err := primitive.ObjectIDFromHex(experimentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid experiment_id format"})
		return
	}

	// Получаем назначения пользователя
	assignments, err := s.assignmentRepo.GetAssignmentsByID("user_id", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get assignments"})
		return
	}

	// Ищем назначение для данного эксперимента
	for _, assignment := range *assignments {
		if assignment.ExperimentID == experimentID {
			c.JSON(http.StatusOK, gin.H{
				"user_id":          userIDStr,
				"experiment_id":    experimentIDStr,
				"assigned_variant": assignment.Variant,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "no assignment found for this user and experiment"})
} 