package gapi

import (
	"context"
	"strings"

	pb "github.com/takeru-a/fakeself_backend/grpc"
	"github.com/takeru-a/fakeself_backend/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (userServer *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error){
	user := models.CreateUserRequest{
		Name: req.GetName(),
		Token: req.GetToken(),
		Answer: 0,
		Score: 0,
	}
	newUser, err := userServer.userService.CreateUser(&user)
	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	
	res := &pb.UserResponse{
		User: &pb.User{
			Id: newUser.Id.Hex(),
			Name: newUser.Name,
			Answer: newUser.Answer,
			Score: newUser.Score,
			CreatedAt: timestamppb.New(newUser.CreateAt),
		},
	}
	return res, nil
}