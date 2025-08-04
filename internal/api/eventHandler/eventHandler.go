package eventHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/nedokyrill/ab-platform/internal/services"
)

type EventHandler struct {
	EventService services.EventServiceInterface
}

func NewEventHandler(EventService services.EventServiceInterface) *EventHandler {
	return &EventHandler{
		EventService: EventService,
	}
}

func (h *EventHandler) InitEventHandler(router *gin.RouterGroup) {
	eventRouter := router.Group("/experiment")
	{
		// POST /api/v1/experiment/event - зафиксировать событие пользователя
		eventRouter.POST("/event", h.EventService.CreateEvent)
		
		// GET /api/v1/experiment/event/stats - получить статистику событий
		eventRouter.GET("/event/stats", h.EventService.GetEventStats)
	}
} 