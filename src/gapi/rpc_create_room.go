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

func (roomServer *RoomServer) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.RoomResponse, error){
	room := models.CreateRoomRequest{
		Host: req.GetHost(),
		Token: req.GetToken(),
		Players: req.GetPlayers(),
		Topic: req.GetTopic(),
	}
	newRoom, err := roomServer.roomService.CreateRoom(&room)
	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.RoomResponse{
		Room: &pb.Room{
			Id: newRoom.Id.Hex(),
			Host: newRoom.Host,
			Token: newRoom.Token,
			Players: newRoom.Players,
			Topic: newRoom.Topic,
			CreatedAt: timestamppb.New(newRoom.CreateAt),
			UpdatedAt: timestamppb.New(newRoom.UpdatedAt),
		},
	}
	return res, nil
}