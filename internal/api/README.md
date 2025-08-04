# API Handlers Layer

Этот пакет содержит HTTP хендлеры для обработки запросов к API.

## Структура хендлеров

Каждый хендлер организован в отдельную директорию с основным файлом:

### assignmentHandler/
- `assignmentHandler.go` - структура, конструктор и регистрация маршрутов для назначения вариантов

### eventHandler/
- `eventHandler.go` - структура, конструктор и регистрация маршрутов для событий пользователей

### experimentHandler/
- `experimentHandler.go` - структура, конструктор и регистрация маршрутов для экспериментов

### userHandler/
- `userHandler.go` - структура, конструктор и регистрация маршрутов для пользователей

### healthHandler/
- `healthHandler.go` - структура, конструктор и регистрация маршрутов для health check и метрик

## API Endpoints

### AssignmentHandler
- `POST /api/v1/experiment/assign` - назначить пользователю вариант A или B
- `GET /api/v1/experiment/variant` - получить назначенный пользователю вариант

### EventHandler
- `POST /api/v1/experiment/event` - зафиксировать событие пользователя
- `GET /api/v1/experiment/event/stats` - получить статистику событий

### ExperimentHandler
- `POST /api/v1/experiments` - создать новый эксперимент
- `GET /api/v1/experiments` - получить список всех экспериментов
- `GET /api/v1/experiments/:id` - получить эксперимент по ID

### UserHandler
- `POST /api/v1/users` - создать нового пользователя
- `GET /api/v1/users` - получить список всех пользователей
- `GET /api/v1/users/:id` - получить пользователя по ID

### HealthHandler
- `GET /api/v1/health` - healthcheck для мониторинга
- `GET /metrics` - метрики Prometheus

## Использование

### Инициализация хендлеров:

```go
// Создание хендлеров
assignmentHandler := assignmentHandler.NewAssignmentHandler(assignmentService)
eventHandler := eventHandler.NewEventHandler(eventService)
experimentHandler := experimentHandler.NewExperimentHandler(experimentService)
userHandler := userHandler.NewUserHandler(userService)
healthHandler := healthHandler.NewHealthHandler(healthService)

// Регистрация маршрутов
v1 := router.Group("/api/v1")
assignmentHandler.InitAssignmentHandlers(v1)
eventHandler.InitEventHandler(v1)
experimentHandler.InitExperimentHandlers(v1)
userHandler.InitUserHandlers(v1)
healthHandler.InitHealthHandlers(v1)
```

### Структура хендлера:

```go
type HandlerName struct {
    ServiceName services.ServiceInterface
}

func NewHandlerName(ServiceName services.ServiceInterface) *HandlerName {
    return &HandlerName{
        ServiceName: ServiceName,
    }
}

func (h *HandlerName) InitHandlerName(router *gin.RouterGroup) {
    handlerRouter := router.Group("/path")
    {
        handlerRouter.POST("/endpoint", h.ServiceName.MethodName)
        handlerRouter.GET("/endpoint", h.ServiceName.MethodName)
    }
}
```

## Особенности

- **Прямое делегирование** - хендлеры напрямую вызывают методы сервисов
- **Упрощенная структура** - нет промежуточных методов
- **Маршрутизация** - каждый хендлер регистрирует свои маршруты
- **Типизация** - использование интерфейсов сервисов для зависимости

## Структура файлов

```
internal/api/
├── README.md                      # Документация API
├── assignmentHandler/
│   └── assignmentHandler.go       # Хендлер назначений
├── eventHandler/
│   └── eventHandler.go            # Хендлер событий
├── experimentHandler/
│   └── experimentHandler.go       # Хендлер экспериментов
├── userHandler/
│   └── userHandler.go             # Хендлер пользователей
└── healthHandler/
    └── healthHandler.go           # Хендлер health check
``` 