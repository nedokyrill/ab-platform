package models

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type EventModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserID       uuid.UUID          `bson:"user_id"`
	ExperimentID uuid.UUID          `bson:"experiment_id"`
	Variant      string             `bson:"variant"`    // "A" или "B"
	EventType    string             `bson:"event_type"` // например: "click", "conversion", "view"
	Timestamp    time.Time          `bson:"timestamp"`
}
