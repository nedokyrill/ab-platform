package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
	"time"
)

var (
	// HTTP метрики
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// A/B тестирование метрики
	experimentAssignments = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "experiment_assignments_total",
			Help: "Total number of experiment assignments by variant",
		},
		[]string{"experiment_id", "variant"},
	)

	// Системные метрики
	memoryUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "memory_usage_bytes",
			Help: "Current memory usage in bytes",
		},
	)

	// RPS (Requests Per Second)
	requestsPerSecond = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "requests_per_second",
			Help: "Requests per second by endpoint",
		},
		[]string{"endpoint"},
	)
)

// InitMetrics инициализирует все метрики
func InitMetrics() {
	prometheus.MustRegister(
		httpRequestsTotal,
		httpRequestDuration,
		experimentAssignments,
		memoryUsage,
		requestsPerSecond,
	)
}

// MetricsMiddleware middleware для сбора HTTP метрик
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Обработка запроса
		c.Next()

		// Сбор метрик
		duration := time.Since(start).Seconds()
		status := c.Writer.Status()
		method := c.Request.Method
		endpoint := c.FullPath()

		// HTTP метрики
		httpRequestsTotal.WithLabelValues(method, endpoint, strconv.Itoa(status)).Inc()
		httpRequestDuration.WithLabelValues(method, endpoint).Observe(duration)

		// RPS метрика (упрощенная версия)
		requestsPerSecond.WithLabelValues(endpoint).Inc()
	}
}

// RecordExperimentAssignment записывает метрику назначения варианта
func RecordExperimentAssignment(experimentID, variant string) {
	experimentAssignments.WithLabelValues(experimentID, variant).Inc()
}

// UpdateMemoryUsage обновляет метрику использования памяти
func UpdateMemoryUsage(bytes uint64) {
	memoryUsage.Set(float64(bytes))
}

// MetricsHandler возвращает HTTP handler для Prometheus
func MetricsHandler() http.Handler {
	return promhttp.Handler()
} 