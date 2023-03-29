package services

import "github.com/takeru-a/fakeself_backend/models"

type UserService interface {
	CreateUser(*models.CreateUserRequest) (*models.User, error)
	UpdateUser(string, *models.UpdateUser) (*models.User, error)
	FindUserById(string) (*models.User, error)
}