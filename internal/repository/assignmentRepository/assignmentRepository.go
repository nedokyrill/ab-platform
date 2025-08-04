package assignmentRepository

import (
	"context"
	"github.com/google/uuid"
	"github.com/nedokyrill/ab-platform/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AssignmentRepository struct {
	db *mongo.Client
}

func NewAssignmentRepository(db *mongo.Client) *AssignmentRepository {
	return &AssignmentRepository{
		db: db,
	}
}

func (r *AssignmentRepository) CreateAssignment(newAssignment models.AssignmentModel) error {
	assignmentCollection := r.db.Database("ab_platform").Collection("assignments")

	_, err := assignmentCollection.InsertOne(context.Background(), newAssignment)
	if err != nil {
		return err
	}
	return nil
}

func (r *AssignmentRepository) GetAssignmentsByID(name string, ID uuid.UUID) (*[]models.AssignmentModel, error) {
	var assignments []models.AssignmentModel
	assignmentCollection := r.db.Database("ab_platform").Collection("assignments")

	cursor, err := assignmentCollection.Find(context.Background(), bson.D{{name, ID}})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var assignment models.AssignmentModel
		if err = cursor.Decode(&assignment); err != nil {
			return nil, err
		}
		assignments = append(assignments, assignment)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return &assignments, nil
}
