package app

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nedokyrill/ab-platform/internal/api/assignmentHandler"
	"github.com/nedokyrill/ab-platform/internal/api/eventHandler"
	"github.com/nedokyrill/ab-platform/internal/api/experimentHandler"
	"github.com/nedokyrill/ab-platform/internal/api/healthHandler"
	"github.com/nedokyrill/ab-platform/internal/api/userHandler"
	"github.com/nedokyrill/ab-platform/internal/metrics"
	"github.com/nedokyrill/ab-platform/internal/metrics/middleware"
	"github.com/nedokyrill/ab-platform/internal/repository/assignmentRepository"
	"github.com/nedokyrill/ab-platform/internal/repository/eventRepository"
	"github.com/nedokyrill/ab-platform/internal/repository/experimentRepository"
	"github.com/nedokyrill/ab-platform/internal/repository/userRepository"
	"github.com/nedokyrill/ab-platform/internal/server"
	"github.com/nedokyrill/ab-platform/internal/services/assignmentService"
	"github.com/nedokyrill/ab-platform/internal/services/eventService"
	"github.com/nedokyrill/ab-platform/internal/services/experimentService"
	"github.com/nedokyrill/ab-platform/internal/services/healthService"
	"github.com/nedokyrill/ab-platform/internal/services/userService"
	"github.com/nedokyrill/ab-platform/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	logger.InitLogger()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Инициализируем метрики
	metrics.InitMetrics()

	// Запускаем сбор системных метрик
	middleware.StartMetricsCollector()

	// MongoDB подключение
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	log.Println(mongoURI)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		logger.Logger.Fatalf("Unable to connect to MongoDB: %v\n", err)
		os.Exit(1)
	}
	logger.Logger.Info("Connected to MongoDB successfully")

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			logger.Logger.Fatalf("Error disconnecting from MongoDB: %v\n", err)
		}
	}()

	// Проверка соединения
	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Logger.Fatalf("MongoDB ping failed: %v\n", err)
	}
	logger.Logger.Info("MongoDB ping successful")

	// Репозитории
	assignmentRepo := assignmentRepository.NewAssignmentRepository(client)
	eventRepo := eventRepository.NewEventRepository(client)
	experimentRepo := experimentRepository.NewExperimentRepository(client)
	userRepo := userRepository.NewUserRepository(client)

	// Сервисы
	assignmentServ := assignmentService.NewAssignmentService(assignmentRepo, experimentRepo, userRepo)
	eventServ := eventService.NewEventService(eventRepo, assignmentRepo, experimentRepo)
	experimentServ := experimentService.NewExperimentService(experimentRepo)
	userServ := userService.NewUserService(userRepo)
	healthServ := healthService.NewHealthService()

	// Хендлеры
	assignmentHandler := assignmentHandler.NewAssignmentHandler(assignmentServ)
	eventHandler := eventHandler.NewEventHandler(eventServ)
	experimentHandler := experimentHandler.NewExperimentHandler(experimentServ)
	userHandler := userHandler.NewUserHandler(userServ)
	healthHandler := healthHandler.NewHealthHandler(healthServ)

	// Настройка роутера
	router := gin.Default()

	// Middleware для метрик
	router.Use(metrics.MetricsMiddleware())
	router.Use(middleware.SystemMetricsMiddleware())

	// CORS настройки
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Настрой под свои нужды
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// GET /metrics - отдаёт метрики Prometheus (RPS, latency, ошибки)
	router.GET("/metrics", gin.WrapH(metrics.MetricsHandler()))

	// API группа
	API := router.Group("/api/v1")

	// Регистрация маршрутов
	assignmentHandler.InitAssignmentHandlers(API)
	eventHandler.InitEventHandler(API)
	experimentHandler.InitExperimentHandlers(API)
	userHandler.InitUserHandlers(API)
	healthHandler.InitHealthHandlers(API)

	// Инициализация и конфигурация HTTP сервера
	srv := server.NewAPIServer(router)

	// Старт сервера
	go srv.Start()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		logger.Logger.Fatalw("Shutdown error", "error", err)
	}
}
