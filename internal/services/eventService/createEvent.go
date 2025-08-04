package eventService

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nedokyrill/ab-platform/internal/models"
	"net/http"
	"time"
)

// CreateEvent создает событие пользователя в эксперименте
func (s *EventService) CreateEvent(c *gin.Context) {
	var request struct {
		UserID       string    `json:"user_id" binding:"required"`
		ExperimentID string    `json:"experiment_id" binding:"required"`
		Event        string    `json:"event" binding:"required"`
		Timestamp    time.Time `json:"timestamp"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	// Парсим user_id
	userID, err := uuid.Parse(request.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id format"})
		return
	}

	// Парсим experiment_id
	experimentID, err := uuid.Parse(request.ExperimentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid experiment_id format"})
		return
	}

	// Проверяем, есть ли назначение для этого пользователя в эксперименте
	assignments, err := s.assignmentRepo.GetAssignmentsByID("user_id", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check user assignment"})
		return
	}

	// Ищем назначение для данного эксперимента
	var assignedVariant string
	found := false
	for _, assignment := range *assignments {
		// Сравниваем experiment_id (нужно преобразовать в UUID для сравнения)
		if assignment.ExperimentID.Hex() == experimentID.String() {
			assignedVariant = assignment.Variant
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user is not assigned to this experiment"})
		return
	}

	// Устанавливаем timestamp если не передан
	if request.Timestamp.IsZero() {
		request.Timestamp = time.Now()
	}

	// Создаем событие
	event := models.EventModel{
		UserID:       userID,
		ExperimentID: experimentID,
		Variant:      assignedVariant,
		EventType:    request.Event,
		Timestamp:    request.Timestamp,
	}

	err = s.eventRepo.CreateEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
} 