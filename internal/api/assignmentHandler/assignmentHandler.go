package assignmentHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/nedokyrill/ab-platform/internal/services"
)

type AssignmentHandler struct {
	AssignmentService services.AssignmentServiceInterface
}

func NewAssignmentHandler(AssignmentService services.AssignmentServiceInterface) *AssignmentHandler {
	return &AssignmentHandler{
		AssignmentService: AssignmentService,
	}
}

func (h *AssignmentHandler) InitAssignmentHandlers(router *gin.RouterGroup) {
	assignmentRouter := router.Group("/experiment")
	{
		// POST /api/v1/experiment/assign - назначить пользователю вариант A или B
		assignmentRouter.POST("/assign", h.AssignmentService.AssignVariant)
		
		// GET /api/v1/experiment/variant - получить назначенный пользователю вариант
		assignmentRouter.GET("/variant", h.AssignmentService.GetAssignedVariant)
	}
} 