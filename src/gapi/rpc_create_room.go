package gapi

import (
	"context"
	"strings"
	"math/rand"
	pb "github.com/takeru-a/fakeself_backend/grpc"
	"github.com/takeru-a/fakeself_backend/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)
func GenerateToken(n int) string {
    var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func (roomServer *RoomServer) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.RoomResponse, error){
	room := models.CreateRoomRequest{
		HostName: req.GetHost(),
		Token: GenerateToken(32),
		Topic: []string{"出身地は？", "趣味は？","好きな食べ物は？","特技は？","昔、やっていた部活は？"},
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
		},
	}
	return res, nil
}