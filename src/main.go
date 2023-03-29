package main

import (
	"context"
	"fmt"
	"log"
	"net"
	// "net/http"

	// "github.com/gin-contrib/cors"
	// "github.com/gin-gonic/gin"
	// "github.com/go-redis/redis/v8"

	"github.com/takeru-a/fakeself_backend/configs"
	"github.com/takeru-a/fakeself_backend/gapi"
	roomservices "github.com/takeru-a/fakeself_backend/services/room"
	userservices "github.com/takeru-a/fakeself_backend/services/user"
	pb "github.com/takeru-a/fakeself_backend/grpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	ctx         context.Context
	mongoclient *mongo.Client
	userService         userservices.UserService
	roomCollection      *mongo.Collection
	userCollection      *mongo.Collection
	roomService         roomservices.RoomService
)

func init() {
	ctx = context.TODO()
	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(configs.EnvMongoURI())
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")


	// Collections
	userCollection = mongoclient.Database("fakeself").Collection("users")
	userService = userservices.NewUserService(userCollection, ctx)
	

	roomCollection = mongoclient.Database("fakeself").Collection("room")
	roomService = roomservices.NewRoomService(roomCollection, ctx)

	// server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)
	startGrpcServer()
}

func startGrpcServer() {
	
	userServer, err := gapi.NewGrpcUserServer(userCollection, userService)
	if err != nil {
		log.Fatal("cannot create grpc userServer: ", err)
	}

	roomServer, err := gapi.NewGrpcRoomServer(roomCollection, roomService)
	if err != nil {
		log.Fatal("cannot create grpc roomServer: ", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServer)
	pb.RegisterRoomServiceServer(grpcServer, roomServer)
	reflection.Register(grpcServer)
	
	listener, err := net.Listen("tcp", configs.GrpcServerAddress())
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
}
