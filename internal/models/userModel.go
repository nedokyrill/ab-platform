package models

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ExternalID uuid.UUID          `bson:"externalId"`
	CreatedAt  time.Time          `bson:"createdAt"`
}
