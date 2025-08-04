package models

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ExperimentModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ExternalID  uuid.UUID          `bson:"externalId"`
	Name        string             `bson:"name"`        // Название эксперимента
	Description string             `bson:"description"` // Описание, что тестируем
	VariantA    string             `bson:"variantA"`    // Имя или ID варианта A
	VariantB    string             `bson:"variantB"`    // Имя или ID варианта B
	CreatedAt   time.Time          `bson:"createdAt"`
}
