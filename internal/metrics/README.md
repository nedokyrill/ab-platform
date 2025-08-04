# Metrics System

Система метрик для A/B тестирования платформы с интеграцией Prometheus и Grafana.

## Собранные метрики

### HTTP метрики (автоматически)
- `http_requests_total` - общее количество запросов по endpoint'ам
- `http_request_duration_seconds` - время ответа сервиса
- `requests_per_second` - запросы в секунду по endpoint'ам

### A/B тестирование метрики
- `experiment_assignments_total` - количество назначений вариантов A/B

### Системные метрики
- `memory_usage_bytes` - использование памяти

## Структура файлов

```
internal/metrics/
├── metrics.go                    # Основные метрики Prometheus
├── middleware/
│   └── system_metrics.go        # Middleware для системных метрик
└── README.md                    # Документация
```

## Использование

### 1. Инициализация метрик

```go
// В main.go или server.go
import "github.com/nedokyrill/ab-platform/internal/metrics"

func main() {
    // Инициализируем метрики
    metrics.InitMetrics()
    
    // Запускаем сбор системных метрик
    middleware.StartMetricsCollector()
    
    // ... остальной код
}
```

### 2. Подключение middleware

```go
// В server.go
import (
    "github.com/nedokyrill/ab-platform/internal/metrics"
    middleware "github.com/nedokyrill/ab-platform/internal/metrics/middleware"
)

func setupRouter() *gin.Engine {
    router := gin.New()
    
    // Подключаем middleware для метрик
    router.Use(metrics.MetricsMiddleware())
    router.Use(middleware.SystemMetricsMiddleware())
    
    // ... остальные middleware и маршруты
}
```

### 3. Запись специфичных метрик

```go
// В сервисах для A/B тестирования
import "github.com/nedokyrill/ab-platform/internal/metrics"

// При назначении варианта
metrics.RecordExperimentAssignment(experimentID, variant)
```

## Prometheus конфигурация

Файл `prometheus/prometheus.yaml` настроен для сбора метрик с:
- `localhost:8080` - основной сервис
- `localhost:9090` - сам Prometheus

## Доступные метрики для Grafana

### Запросы для Grafana:

```promql
# Общее количество запросов
http_requests_total

# Время ответа (среднее)
rate(http_request_duration_seconds_sum[5m]) / rate(http_request_duration_seconds_count[5m])

# Запросы в секунду
rate(http_requests_total[5m])

# Процент ошибок
rate(http_requests_total{status=~"4..|5.."}[5m]) / rate(http_requests_total[5m]) * 100

# Назначения вариантов
experiment_assignments_total

# Использование памяти
memory_usage_bytes
```

## Endpoints

- `GET /metrics` - метрики Prometheus
- `GET /api/v1/health` - health check

## Интеграция с Grafana

1. **Добавить Prometheus как Data Source**
2. **Создать дашборды** с запросами выше
3. **Настроить алерты** для критических метрик

## Примеры дашбордов

### Системный мониторинг:
- RPS по endpoint'ам
- Время ответа
- Процент ошибок
- Использование памяти

### A/B тестирование:
- Количество назначений по вариантам
- Распределение A/B (должно быть 50/50)
- Активность экспериментов 