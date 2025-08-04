package experimentRepository

import (
	"context"
	"github.com/nedokyrill/ab-platform/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExperimentRepository struct {
	db *mongo.Client
}

func NewExperimentRepository(db *mongo.Client) *ExperimentRepository {
	return &ExperimentRepository{
		db: db,
	}
}

func (r *ExperimentRepository) CreateExperiment(newExperiment models.ExperimentModel) error {
	experimentCollection := r.db.Database("ab_platform").Collection("experiments")

	_, err := experimentCollection.InsertOne(context.Background(), newExperiment)
	if err != nil {
		return err
	}

	return nil
}

func (r *ExperimentRepository) GetAllExperiments() (*[]models.ExperimentModel, error) {
	var experiments []models.ExperimentModel
	experimentCollection := r.db.Database("ab_platform").Collection("experiments")

	cursor, err := experimentCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var experiment models.ExperimentModel
		err = cursor.Decode(&experiment)
		if err != nil {
			return nil, err
		}
		experiments = append(experiments, experiment)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}
	return &experiments, nil
}
