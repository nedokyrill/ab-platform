package repository

import (
	"github.com/google/uuid"
	"github.com/nedokyrill/ab-platform/internal/models"
)

type UserRepositoryInterface interface {
	CreateUser(newUser models.UserModel) error
	GetAllUsers() (*[]models.UserModel, error)
}

type AssignmentRepositoryInterface interface {
	CreateAssignment(newAssignment models.AssignmentModel) error
	GetAssignmentsByID(name string, ID uuid.UUID) (*[]models.AssignmentModel, error)
}

type EventRepositoryInterface interface {
	CreateEvent(newEvent models.EventModel) error
	GetEventsByID(name string, ID uuid.UUID) (*[]models.EventModel, error)
}

type ExperimentRepositoryInterface interface {
	CreateExperiment(newExperiment models.ExperimentModel) error
	GetAllExperiments() (*[]models.ExperimentModel, error)
}
