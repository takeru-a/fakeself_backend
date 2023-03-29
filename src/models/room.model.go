package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	pb "github.com/takeru-a/fakeself_backend/grpc"
)

type CreateRoomRequest struct {
	Host    *pb.User    `json:"host" bson:"host" binding:"required"`
	Token   string    `json:"token" bson:"token" binding:"required"`
	Players   []*pb.User    `json:"players,omitempty" bson:"players,omitempty"`
	Topic   []string    `json:"topic" bson:"topic" binding:"required"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Room struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Host    *pb.User    `json:"host" bson:"host" binding:"required"`
	Token   string    `json:"token" bson:"token" binding:"required"`
	Players   []*pb.User    `json:"players,omitempty" bson:"players,omitempty"`
	Topic   []string    `json:"topic" bson:"topic" binding:"required"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateRoom struct {
	Host    *pb.User    `json:"host" bson:"host" binding:"required"`
	Token   string    `json:"token" bson:"token" binding:"required"`
	Players   []*pb.User     `json:"players,omitempty" bson:"players,omitempty"`
	Topic   []string    `json:"topic" bson:"topic" binding:"required"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}