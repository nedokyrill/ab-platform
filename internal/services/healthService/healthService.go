package healthService

type HealthService struct {
	// Можно добавить зависимости для проверки состояния БД и других сервисов
}

func NewHealthService() *HealthService {
	return &HealthService{}
} 