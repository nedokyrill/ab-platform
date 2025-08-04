package assignmentService

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nedokyrill/ab-platform/internal/metrics"
	"github.com/nedokyrill/ab-platform/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

// AssignVariant назначает пользователю вариант A или B в эксперименте
func (s *AssignmentService) AssignVariant(c *gin.Context) {
	userIDStr := c.Query("user_id")
	experimentIDStr := c.Query("experiment_id")

	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id format"})
		return
	}

	// Если experiment_id не передан, используем первый активный эксперимент
	var experimentID primitive.ObjectID
	if experimentIDStr == "" {
		experiments, err := s.experimentRepo.GetAllExperiments()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get experiments"})
			return
		}
		if len(*experiments) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "no active experiments found"})
			return
		}
		experimentID = (*experiments)[0].ID
	} else {
		// Парсим experiment_id если передан
		experimentID, err = primitive.ObjectIDFromHex(experimentIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid experiment_id format"})
			return
		}
	}

	// Проверяем, есть ли уже назначение для этого пользователя в эксперименте
	existingAssignments, err := s.assignmentRepo.GetAssignmentsByID("user_id", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check existing assignments"})
		return
	}

	// Ищем существующее назначение для данного эксперимента
	for _, assignment := range *existingAssignments {
		if assignment.ExperimentID == experimentID {
			c.JSON(http.StatusOK, gin.H{
				"user_id":          userIDStr,
				"experiment_id":    experimentIDStr,
				"assigned_variant": assignment.Variant,
			})
			return
		}
	}

	// Назначаем вариант по бизнес-логике: остаток от деления ID на 2
	variant := s.determineVariant(userID)

	// Создаем новое назначение
	assignment := models.AssignmentModel{
		UserID:       userID,
		ExperimentID: experimentID,
		Variant:      variant,
		AssignedAt:   time.Now(),
	}

	err = s.assignmentRepo.CreateAssignment(assignment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create assignment"})
		return
	}

	// Записываем метрику назначения варианта
	metrics.RecordExperimentAssignment(experimentID.Hex(), variant)

	c.JSON(http.StatusOK, gin.H{
		"user_id":          userIDStr,
		"experiment_id":    experimentIDStr,
		"assigned_variant": variant,
	})
}

// determineVariant определяет вариант по остатку от деления ID пользователя на 2
func (s *AssignmentService) determineVariant(userID uuid.UUID) string {
	bytes := userID[:]
	lastByte := bytes[len(bytes)-1]

	// Если последний байт четный - вариант A, иначе - вариант B
	if lastByte%2 == 0 {
		return "A"
	}
	return "B"
}
