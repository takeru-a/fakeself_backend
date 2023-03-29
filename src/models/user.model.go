package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserRequest struct {
	Name     string    `json:"name" bson:"name" binding:"required"`
	Answer   int    `json:"answer" bson:"answer" binding:"required"`
	Score    int    `json:"score,omitempty" bson:"score,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string    `json:"name" bson:"name" binding:"required"`
	Answer   int    `json:"answer" bson:"answer" binding:"required"`
	Score    int    `json:"score,omitempty" bson:"score,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateUser struct {
	Name     string    `json:"name" bson:"name" binding:"required"`
	Answer   int    `json:"answer" bson:"answer" binding:"required"`
	Score    int    `json:"score,omitempty" bson:"score,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

