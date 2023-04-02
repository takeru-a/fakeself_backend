package gapi

import(
	pb "github.com/takeru-a/fakeself_backend/grpc"
	roomservices "github.com/takeru-a/fakeself_backend/services/room"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomServer struct{
	pb.UnimplementedRoomServiceServer
	DB *mongo.Database
	roomService roomservices.RoomService
}

func NewGrpcRoomServer(DB *mongo.Database, roomService roomservices.RoomService) (*RoomServer, error){
	roomServer := &RoomServer{
		DB: DB,
		roomService: roomService,
	}

	return roomServer, nil
}