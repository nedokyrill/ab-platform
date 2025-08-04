package userRepository

import (
	"context"
	"github.com/nedokyrill/ab-platform/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Client
}

func NewUserRepository(db *mongo.Client) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(newUser models.UserModel) error {
	assignmentCollection := r.db.Database("ab_platform").Collection("users")

	_, err := assignmentCollection.InsertOne(context.Background(), newUser)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetAllUsers() (*[]models.UserModel, error) {
	var users []models.UserModel
	userCollection := r.db.Database("ab_platform").Collection("users")

	cursor, err := userCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user models.UserModel
		if err = cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}
