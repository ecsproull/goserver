package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BlogID         primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	CommenterName  string             `json:"commenter_name" bson:"commenter_name"`
	CommenterEmail string             `json:"commenter_email" bson:"commenter_email"`
	CommentBody    string             `json:"comment_body" bson:"comment_body"`
	CreatedAt      time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt" bson:"updatedAt"`
}
