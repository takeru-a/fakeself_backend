package services

import "github.com/takeru-a/fakeself_backend/models"

type RoomService interface {
	CreateRoom(*models.CreateRoomRequest) (*models.Room, error)
	UpdateRoom(string, *models.UpdateRoom) (*models.Room, error)
	FindRoomById(string) (*models.Room, error)
	FindRoomByToken(string) (*models.Room, error)
}