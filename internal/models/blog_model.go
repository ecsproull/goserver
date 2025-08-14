package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Subject    string             `json:"blog_subject" bson:"blog_subject"`
	OwnerName  string             `json:"blog_owner_name" bson:"blog_owner_name"`
	OwnerEmail string             `json:"blog_owner_email" bson:"blog_owner_email"`
	Body       string             `json:"blog_body" bson:"blog_body"`
	Category   string             `json:"blog_category" bson:"blog_category"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt"`
}
