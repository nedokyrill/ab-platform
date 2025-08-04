package eventRepository

import (
	"context"
	"github.com/google/uuid"
	"github.com/nedokyrill/ab-platform/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventRepository struct {
	db *mongo.Client
}

func NewEventRepository(db *mongo.Client) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) CreateEvent(newEvent models.EventModel) error {
	eventsCollection := r.db.Database("ab_platform").Collection("events")

	_, err := eventsCollection.InsertOne(context.Background(), newEvent)
	if err != nil {
		return err
	}

	return nil
}

func (r *EventRepository) GetEventsByID(name string, ID uuid.UUID) (*[]models.EventModel, error) {
	var events []models.EventModel
	eventsCollection := r.db.Database("ab_platform").Collection("events")

	cursor, err := eventsCollection.Find(context.Background(), bson.D{{name, ID}})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var event models.EventModel
		if err = cursor.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return &events, nil
}
