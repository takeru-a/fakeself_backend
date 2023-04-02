package gapi

import (
	pb "github.com/takeru-a/fakeself_backend/grpc"
	userservices "github.com/takeru-a/fakeself_backend/services/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServer struct{
	pb.UnimplementedUserServiceServer
	DB *mongo.Database
	userService userservices.UserService
}

func NewGrpcUserServer(DB *mongo.Database, userService userservices.UserService) (*UserServer, error){
	userServer := &UserServer{
		DB: DB,
		userService: userService,
	}

	return userServer, nil
}