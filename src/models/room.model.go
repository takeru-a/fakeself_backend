package models

import (
	"time"
	pb "github.com/takeru-a/fakeself_backend/grpc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateRoomRequest struct {
	HostName string    `json:"host" bson:"host" binding:"required"`
	Token   string    `json:"token" bson:"token" binding:"required"`
	Topic   []string    `json:"topic" bson:"topic" binding:"required"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`

}

type DBRoom struct {
	Host    *User    `json:"host" bson:"host" binding:"required"`
	Token   string    `json:"token" bson:"token" binding:"required"`
	Players   []*User    `json:"players,omitempty" bson:"players,omitempty"`
	Topic   []string    `json:"topic" bson:"topic" binding:"required"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	
}

type Room struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Host    *pb.User    `json:"host" bson:"host" binding:"required"`
	Token   string    `json:"token" bson:"token" binding:"required"`
	Players   []*pb.User    `json:"players,omitempty" bson:"players,omitempty"`
	Topic   []string    `json:"topic" bson:"topic" binding:"required"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	
}

type UpdateRoom struct {
	Token   string    `json:"token" bson:"token" binding:"required"`
	Players   []*pb.User     `json:"players,omitempty" bson:"players,omitempty"`
	Topic   []string    `json:"topic" bson:"topic" binding:"required"`
}