package gapi

import (
	pb "github.com/takeru-a/fakeself_backend/grpc"
	userservices "github.com/takeru-a/fakeself_backend/services/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServer struct{
	pb.UnimplementedUserServiceServer
	userCollection *mongo.Collection
	userService userservices.UserService
}

func NewGrpcUserServer(userCollection *mongo.Collection, userService userservices.UserService) (*UserServer, error){
	userServer := &UserServer{
		userCollection: userCollection,
		userService: userService,
	}

	return userServer, nil
}