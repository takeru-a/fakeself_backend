package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserRequest struct {
	Name     string    `json:"name" bson:"name" binding:"required"`
	Token   string    `json:"token" bson:"token" binding:"required"`
	Answer   int64    `json:"answer,omitempty" bson:"answer,omitempty"`
	Score    int64    `json:"score,omitempty" bson:"score,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	
}

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string    `json:"name,omitempty" bson:"name,omitempty" `
	Answer   int64    `json:"answer" bson:"answer" binding:"required"`
	Score    int64    `json:"score,omitempty" bson:"score,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`

}

type DBUser struct {
	Name     string    `json:"name,omitempty" bson:"name,omitempty" `
	Answer   int64    `json:"answer" bson:"answer" binding:"required"`
	Score    int64    `json:"score,omitempty" bson:"score,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type UpdateUser struct {
	Name     string    `json:"name" bson:"name" binding:"required"`
	Answer   int64    `json:"answer,omitempty" bson:"answer,omitempty" `
	Score    int64    `json:"score,omitempty" bson:"score,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`

}

