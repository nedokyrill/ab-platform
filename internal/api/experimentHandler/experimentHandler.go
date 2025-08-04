package experimentHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/nedokyrill/ab-platform/internal/services"
)

type ExperimentHandler struct {
	ExperimentService services.ExperimentServiceInterface
}

func NewExperimentHandler(ExperimentService services.ExperimentServiceInterface) *ExperimentHandler {
	return &ExperimentHandler{
		ExperimentService: ExperimentService,
	}
}

func (h *ExperimentHandler) InitExperimentHandlers(router *gin.RouterGroup) {
	experimentRouter := router.Group("/experiments")
	{
		// POST /api/v1/experiments - создать новый эксперимент
		experimentRouter.POST("/", h.ExperimentService.CreateExperiment)
		
		// GET /api/v1/experiments - получить список всех экспериментов
		experimentRouter.GET("/", h.ExperimentService.GetExperiments)
		
		// GET /api/v1/experiments/:id - получить эксперимент по ID
		experimentRouter.GET("/:id", h.ExperimentService.GetExperimentByID)
	}
}
