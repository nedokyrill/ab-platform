package eventService

import (
	"github.com/nedokyrill/ab-platform/internal/repository"
)

type EventService struct {
	eventRepo       repository.EventRepositoryInterface
	assignmentRepo  repository.AssignmentRepositoryInterface
	experimentRepo  repository.ExperimentRepositoryInterface
}

func NewEventService(
	eventRepo repository.EventRepositoryInterface,
	assignmentRepo repository.AssignmentRepositoryInterface,
	experimentRepo repository.ExperimentRepositoryInterface,
) *EventService {
	return &EventService{
		eventRepo:      eventRepo,
		assignmentRepo: assignmentRepo,
		experimentRepo: experimentRepo,
	}
} 