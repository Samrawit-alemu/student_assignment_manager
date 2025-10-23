package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Assignment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	DueDate     primitive.DateTime `bson:"due_date" json:"due_date"`
	Done        bool               `bson:"done" json:"done"`
}
