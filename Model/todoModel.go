package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id             primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Title          string               `bson:"title,omitempty" json:"title,omitempty"`
	Description    string               `bson:"description,omitempty" json:"description,omitempty"`
}