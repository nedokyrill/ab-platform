package experimentService

import (
	"github.com/nedokyrill/ab-platform/internal/repository"
)

type ExperimentService struct {
	experimentRepo repository.ExperimentRepositoryInterface
}

func NewExperimentService(experimentRepo repository.ExperimentRepositoryInterface) *ExperimentService {
	return &ExperimentService{
		experimentRepo: experimentRepo,
	}
} 