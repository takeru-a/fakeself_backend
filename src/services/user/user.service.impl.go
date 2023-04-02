package services

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/takeru-a/fakeself_backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserServiceImpl struct {
	DB *mongo.Database
	ctx            context.Context
}

func NewUserService(DB *mongo.Database, ctx context.Context) UserService {
	return &UserServiceImpl{DB, ctx}
}

func (ur *UserServiceImpl) CreateUser(user *models.CreateUserRequest) (*models.User,error){
	userCollection := ur.DB.Collection("users")
	roomCollection := ur.DB.Collection("room")
	user.CreateAt = time.Now()
	res, err := userCollection.InsertOne(ur.ctx, &models.DBUser{
		Name: user.Name,
		Answer: user.Answer,
		Score: user.Score,
		CreateAt: user.CreateAt,
	})
	if err != nil{
		return nil, err
	}
	var newUser *models.User
	query := bson.M{"_id": res.InsertedID}
	
	if err = userCollection.FindOne(ur.ctx, query).Decode(&newUser); err != nil {
		return nil, err
	}
	var updateRoom *models.Room
	err = roomCollection.FindOne(ur.ctx, bson.M{"token": user.Token}).Decode(&updateRoom)
    if err != nil {
        return nil, err
    }
	log.Printf("%s", updateRoom.Id)
	update := bson.M{"players": newUser}
	_, err = roomCollection.UpdateOne(ur.ctx,
		bson.M{"_id": updateRoom.Id},
        bson.M{
        "$push": update,
        },
	)
	if err != nil {
        return nil, err
    }

	return newUser, nil
}

func (ur *UserServiceImpl) UpdateUser(id string, data *models.UpdateUser) (*models.User, error) {
	userCollection := ur.DB.Collection("users")
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: data}}
	res := userCollection.FindOneAndUpdate(ur.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedUser *models.User
	if err := res.Decode(&updatedUser); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatedUser, nil
}

func (ur *UserServiceImpl) FindUserById(id string) (*models.User, error) {
	userCollection := ur.DB.Collection("users")
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var user *models.User

	if err := userCollection.FindOne(ur.ctx, query).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return user, nil
}

