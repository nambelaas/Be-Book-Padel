package auth

import (
	"Be-Book-Padel/models"
	refreshtoken "Be-Book-Padel/models/repository/refresh_token"
	"Be-Book-Padel/models/repository/user"
)

type AuthServiceInterface interface {
	Register(user *models.Users) (string, string, error)
	Login(email, password string) (string, string, error)
}

type NewAuthService struct {
	UserRepo         user.UserRepository
	RefreshTokenRepo refreshtoken.RefreshTokenRepository
}
