package eventService

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nedokyrill/ab-platform/internal/models"
	"net/http"
)

// GetEventStats получает статистику событий для эксперимента
func (s *EventService) GetEventStats(c *gin.Context) {
	experimentIDStr := c.Query("experiment_id")
	eventType := c.Query("event_type")

	if experimentIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "experiment_id is required"})
		return
	}

	experimentID, err := uuid.Parse(experimentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid experiment_id format"})
		return
	}

	// Получаем все события для эксперимента
	events, err := s.eventRepo.GetEventsByID("experiment_id", experimentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get events"})
		return
	}

	// Фильтруем по типу события если указан
	var filteredEvents []models.EventModel
	if eventType != "" {
		for _, event := range *events {
			if event.EventType == eventType {
				filteredEvents = append(filteredEvents, event)
			}
		}
	} else {
		filteredEvents = *events
	}

	// Подсчитываем статистику по вариантам
	stats := make(map[string]int)
	for _, event := range filteredEvents {
		stats[event.Variant]++
	}

	c.JSON(http.StatusOK, gin.H{
		"experiment_id": experimentIDStr,
		"event_type":    eventType,
		"stats":         stats,
		"total_events":  len(filteredEvents),
	})
}
