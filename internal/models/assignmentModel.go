package models

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AssignmentModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserID       uuid.UUID          `bson:"user_id"` // ExternalID пользователя
	ExperimentID primitive.ObjectID `bson:"experiment_id"`
	Variant      string             `bson:"variant"` // "A" или "B"
	AssignedAt   time.Time          `bson:"assigned_at"`
}
