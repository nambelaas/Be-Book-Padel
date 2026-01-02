package user

import "Be-Book-Padel/models"

type UserRepositoryInterface interface {
	Create(user *models.Users) error
	FindByEmail(email string) (*models.Users, error)
}

type UserRepository struct{}
