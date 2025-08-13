package models

import (
	"time"
)

type Blog struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Subject   string    `json:"blog_subject" bson:"blog_subject"`
	OwnerName string    `json:"blog_owner_name" bson:"blog_owner_name"`
	Body      string    `json:"blog_body" bson:"blog_body"`
	Category  string    `json:"blog_category" bson:"blog_category"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Comment struct {
	ID             string    `json:"_id,omitempty" bson:"_id,omitempty"`
	BlogID         string    `json:"blog_id" bson:"blog_id"`
	CommenterName  string    `json:"commenter_name" bson:"commenter_name"`
	CommenterEmail string    `json:"commenter_email" bson:"commenter_email"`
	CommentBody    string    `json:"comment_body" bson:"comment_body"`
	CreatedAt      time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt" bson:"updatedAt"`
}
