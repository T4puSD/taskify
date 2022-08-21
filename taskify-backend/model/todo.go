package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id      primitive.ObjectID `json:"id" validate:"omitempty" bson:"_id"`
	Content string             `json:"content" validate:"required" bson:"content"`
	IsDone  bool               `json:"is_done" bson:"is_done"`
}
