package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShortURL struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Url     string             `json:"original_url" bson:"original_url,omitempty"`
	ShortId string             `json:"short_id" bson:"short_id,omitempty"`
}
