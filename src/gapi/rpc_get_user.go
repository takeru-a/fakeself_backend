package gapi

import (
	"context"
	"strings"

	pb "github.com/takeru-a/fakeself_backend/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (userServer *UserServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error){
	user, err := userServer.userService.FindUserById(req.GetId())
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())

		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UserResponse{
		User: &pb.User{
			Id: user.Id.Hex(),
			Name: user.Name,
			Answer: user.Answer,
			Score: user.Score,
			CreatedAt: timestamppb.New(user.CreateAt),
		},
	}
	return res, nil
}