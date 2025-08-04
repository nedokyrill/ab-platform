package assignmentService

import (
	"github.com/nedokyrill/ab-platform/internal/repository"
)

type AssignmentService struct {
	assignmentRepo repository.AssignmentRepositoryInterface
	experimentRepo repository.ExperimentRepositoryInterface
	userRepo       repository.UserRepositoryInterface
}

func NewAssignmentService(
	assignmentRepo repository.AssignmentRepositoryInterface,
	experimentRepo repository.ExperimentRepositoryInterface,
	userRepo repository.UserRepositoryInterface,
) *AssignmentService {
	return &AssignmentService{
		assignmentRepo: assignmentRepo,
		experimentRepo: experimentRepo,
		userRepo:       userRepo,
	}
} 