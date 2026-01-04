package user

import (
	"Be-Book-Padel/models"
	"Be-Book-Padel/models/repository/user"
)

type UserServiceInterface interface {
	GetUserByID(id uint) (*models.Users, error)
}

type NewUserService struct{
	UserRepo user.UserRepository
}
