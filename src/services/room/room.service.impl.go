package services

import (
	"context"
	"errors"
	"time"

	"github.com/takeru-a/fakeself_backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RoomServiceImpl struct {
	DB *mongo.Database
	ctx            context.Context
}

func NewRoomService(DB *mongo.Database, ctx context.Context) RoomService {
	return &RoomServiceImpl{DB, ctx}
}

func (rs *RoomServiceImpl) CreateRoom(room *models.CreateRoomRequest) (*models.Room,error){
	roomCollection := rs.DB.Collection("room")
	userCollection := rs.DB.Collection("users")
	hostinfo, err := userCollection.InsertOne(rs.ctx, 
		&models.CreateUserRequest{
			Name: room.HostName,
			Answer: 0,
			Score: 0,
		})
	if err != nil{
		return nil, err
	}
	var newHost *models.User
	query := bson.M{"_id": hostinfo.InsertedID}
	if err = userCollection.FindOne(rs.ctx, query).Decode(&newHost); err != nil {
		return nil, err
	}
	room.CreateAt = time.Now()
	
	input := &models.DBRoom{
		Host: newHost,
		Token: room.Token,
		Players: []*models.User{newHost},
		Topic: room.Topic,
		CreateAt: room.CreateAt,
	}
	res, err := roomCollection.InsertOne(rs.ctx, input)
	if err != nil{
		return nil, err
	}

	var newRoom *models.Room
	query = bson.M{"_id": res.InsertedID}
	if err = roomCollection.FindOne(rs.ctx, query).Decode(&newRoom); err != nil {
		return nil, err
	}

	return newRoom, nil
}

func (rs *RoomServiceImpl) UpdateRoom(id string, data *models.UpdateRoom) (*models.Room, error) {
	roomCollection := rs.DB.Collection("room")
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: data}}
	res := roomCollection.FindOneAndUpdate(rs.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedRoom *models.Room
	if err := res.Decode(&updatedRoom); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatedRoom, nil
}

func (rs *RoomServiceImpl) FindRoomById(id string) (*models.Room, error) {
	roomCollection := rs.DB.Collection("room")
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var room *models.Room

	if err := roomCollection.FindOne(rs.ctx, query).Decode(&room); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return room, nil
}

func (rs *RoomServiceImpl) FindRoomByToken(token string) (*models.Room, error) {
	roomCollection := rs.DB.Collection("room")
	query := bson.M{"token": token}

	var room *models.Room

	if err := roomCollection.FindOne(rs.ctx, query).Decode(&room); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return room, nil
}
