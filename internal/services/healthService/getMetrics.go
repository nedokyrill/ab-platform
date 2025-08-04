package healthService

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetMetrics возвращает метрики для Prometheus
func (s *HealthService) GetMetrics(c *gin.Context) {
	// Здесь можно добавить сбор метрик (RPS, latency, ошибки)
	// Пока возвращаем базовую информацию
	
	c.JSON(http.StatusOK, gin.H{
		"metrics": gin.H{
			"status": "available",
			"note":   "Prometheus metrics endpoint - implement actual metrics collection",
		},
	})
} 