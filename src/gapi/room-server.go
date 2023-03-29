package gapi

import(
	pb "github.com/takeru-a/fakeself_backend/grpc"
	roomservices "github.com/takeru-a/fakeself_backend/services/room"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomServer struct{
	pb.UnimplementedRoomServiceServer
	roomCollection *mongo.Collection
	roomService roomservices.RoomService
}

func NewGrpcRoomServer(roomCollection *mongo.Collection, roomService roomservices.RoomService) (*RoomServer, error){
	roomServer := &RoomServer{
		roomCollection: roomCollection,
		roomService: roomService,
	}

	return roomServer, nil
}