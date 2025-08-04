package healthHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/nedokyrill/ab-platform/internal/services"
)

type HealthHandler struct {
	HealthService services.HealthServiceInterface
}

func NewHealthHandler(HealthService services.HealthServiceInterface) *HealthHandler {
	return &HealthHandler{
		HealthService: HealthService,
	}
}

func (h *HealthHandler) InitHealthHandlers(router *gin.RouterGroup) {
	// GET /api/v1/health - healthcheck для мониторинга
	router.GET("/health", h.HealthService.HealthCheck)
}
