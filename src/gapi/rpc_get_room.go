package gapi

import (
	"context"
	"strings"

	pb "github.com/takeru-a/fakeself_backend/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (roomServer *RoomServer) GetRoomById(ctx context.Context, req *pb.RoomIdRequest) (*pb.RoomResponse, error){
	room, err := roomServer.roomService.FindRoomById(req.GetId())
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())

		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.RoomResponse{
		Room: &pb.Room{
			Id: room.Id.Hex(),
			Host: room.Host,
			Token: room.Token,
			Players: room.Players,
			Topic: room.Topic,
			CreatedAt: timestamppb.New(room.CreateAt),
		},
	}
	return res, nil
}