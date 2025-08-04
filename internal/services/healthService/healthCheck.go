package healthService

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheck проверяет состояние сервиса
func (s *HealthService) HealthCheck(c *gin.Context) {
	// Здесь можно добавить проверки состояния БД, кэша и других зависимостей
	// Пока возвращаем простой статус "healthy"
	
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
} 