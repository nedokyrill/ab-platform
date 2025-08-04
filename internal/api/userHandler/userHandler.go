package userHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/nedokyrill/ab-platform/internal/services"
)

type UserHandler struct {
	UserService services.UserServiceInterface
}

func NewUserHandler(UserService services.UserServiceInterface) *UserHandler {
	return &UserHandler{
		UserService: UserService,
	}
}

func (h *UserHandler) InitUserHandlers(router *gin.RouterGroup) {
	userRouter := router.Group("/users")
	{
		// POST /api/v1/users - создать нового пользователя
		userRouter.POST("/", h.UserService.CreateUser)
		
		// GET /api/v1/users - получить список всех пользователей
		userRouter.GET("/", h.UserService.GetUsers)
		
		// GET /api/v1/users/:id - получить пользователя по ID
		userRouter.GET("/:id", h.UserService.GetUserByID)
	}
} 