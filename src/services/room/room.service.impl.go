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
	roomCollection *mongo.Collection
	ctx            context.Context
}

func NewRoomService(roomCollection *mongo.Collection, ctx context.Context) RoomService {
	return &RoomServiceImpl{roomCollection, ctx}
}

func (rs *RoomServiceImpl) CreateRoom(room *models.CreateRoomRequest) (*models.Room,error){
	room.CreateAt = time.Now()
	room.UpdatedAt = room.CreateAt
	res, err := rs.roomCollection.InsertOne(rs.ctx, room)
	if err != nil{
		return nil, err
	}

	var newRoom *models.Room
	query := bson.M{"_id": res.InsertedID}
	if err = rs.roomCollection.FindOne(rs.ctx, query).Decode(&newRoom); err != nil {
		return nil, err
	}

	return newRoom, nil
}

func (rs *RoomServiceImpl) UpdateRoom(id string, data *models.UpdateRoom) (*models.Room, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: data}}
	res := rs.roomCollection.FindOneAndUpdate(rs.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedRoom *models.Room
	if err := res.Decode(&updatedRoom); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatedRoom, nil
}

func (rs *RoomServiceImpl) FindRoomById(id string) (*models.Room, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var room *models.Room

	if err := rs.roomCollection.FindOne(rs.ctx, query).Decode(&room); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return room, nil
}

func (rs *RoomServiceImpl) FindRoomByToken(token string) (*models.Room, error) {
	
	query := bson.M{"token": token}

	var room *models.Room

	if err := rs.roomCollection.FindOne(rs.ctx, query).Decode(&room); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return room, nil
}
