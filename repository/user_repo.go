package repository

import (
	"context"
	"student_assignment_management/config"
	"student_assignment_management/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct{}

func (r *UserRepository) Collection() *mongo.Collection {
	return config.DB.Collection("users")
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	_, err := r.Collection().InsertOne(context.TODO(), user)
	return err
}

func (r *UserRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.Collection().FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByID(id string) (*entity.User, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	var user entity.User
	err := r.Collection().FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
