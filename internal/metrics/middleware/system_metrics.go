package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nedokyrill/ab-platform/internal/metrics"
	"runtime"
	"time"
)

// SystemMetricsMiddleware собирает системные метрики
func SystemMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	// Обновляем метрики памяти
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		metrics.UpdateMemoryUsage(m.Alloc)

		// Обработка запроса
		c.Next()

		// Логируем ошибки (статус >= 400)
		if c.Writer.Status() >= 400 {
			// Метрика ошибок уже собирается в основном middleware
			// Здесь можно добавить дополнительную логику
		}
	}
}

// StartMetricsCollector запускает периодический сбор метрик
func StartMetricsCollector() {
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			metrics.UpdateMemoryUsage(m.Alloc)
		}
	}()
} 