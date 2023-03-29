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

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{userCollection, ctx}
}

func (ur *UserServiceImpl) CreateUser(user *models.CreateUserRequest) (*models.User,error){
	user.CreateAt = time.Now()
	user.UpdatedAt = user.CreateAt
	res, err := ur.userCollection.InsertOne(ur.ctx, user)
	if err != nil{
		return nil, err
	}

	var newUser *models.User
	query := bson.M{"_id": res.InsertedID}
	if err = ur.userCollection.FindOne(ur.ctx, query).Decode(&newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (ur *UserServiceImpl) UpdateUser(id string, data *models.UpdateUser) (*models.User, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: data}}
	res := ur.userCollection.FindOneAndUpdate(ur.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedUser *models.User
	if err := res.Decode(&updatedUser); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatedUser, nil
}

func (ur *UserServiceImpl) FindUserById(id string) (*models.User, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var user *models.User

	if err := ur.userCollection.FindOne(ur.ctx, query).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return user, nil
}

