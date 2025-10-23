package repository

import (
	"context"
	"student_assignment_management/config"
	"student_assignment_management/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AssignmentRepository struct{}

func (r *AssignmentRepository) Collection() *mongo.Collection {
	return config.DB.Collection("assignments")
}

// Create
func (r *AssignmentRepository) Create(a *entity.Assignment) error {
	_, err := r.Collection().InsertOne(context.TODO(), a)
	return err
}

// Get by ID
func (r *AssignmentRepository) GetByID(id string) (*entity.Assignment, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	var a entity.Assignment
	err := r.Collection().FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// Get all for a user
func (r *AssignmentRepository) GetByUser(userID string) ([]entity.Assignment, error) {
	objID, _ := primitive.ObjectIDFromHex(userID)
	cursor, err := r.Collection().Find(context.TODO(), bson.M{"user_id": objID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var assignments []entity.Assignment
	for cursor.Next(context.TODO()) {
		var a entity.Assignment
		cursor.Decode(&a)
		assignments = append(assignments, a)
	}
	return assignments, nil
}

// Update Done
func (r *AssignmentRepository) UpdateDone(id string, done bool) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.Collection().UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"done": done}},
	)
	return err
}

// Delete
func (r *AssignmentRepository) Delete(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.Collection().DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}
