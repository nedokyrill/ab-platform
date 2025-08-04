package services

import "github.com/gin-gonic/gin"

// ExperimentServiceInterface интерфейс для работы с экспериментами
type ExperimentServiceInterface interface {
	CreateExperiment(c *gin.Context)
	GetExperiments(c *gin.Context)
	GetExperimentByID(c *gin.Context)
}

// AssignmentServiceInterface интерфейс для работы с назначениями вариантов
type AssignmentServiceInterface interface {
	AssignVariant(c *gin.Context)
	GetAssignedVariant(c *gin.Context)
}

// EventServiceInterface интерфейс для работы с событиями
type EventServiceInterface interface {
	CreateEvent(c *gin.Context)
	GetEventStats(c *gin.Context)
}

// UserServiceInterface интерфейс для работы с пользователями
type UserServiceInterface interface {
	CreateUser(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
}

// HealthServiceInterface интерфейс для health check и метрик
type HealthServiceInterface interface {
	HealthCheck(c *gin.Context)
	GetMetrics(c *gin.Context)
}
