package user

import (
	"Be-Book-Padel/database"
	"Be-Book-Padel/models"
)

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (r *UserRepository) Create(user *models.Users) error {
	result := database.DB.Create(user)

	return result.Error
}

func (r *UserRepository) FindByEmail(email string) *models.Users {
	var user models.Users
	result := database.DB.Where("email = ?", email).First(&user)
	if result.RowsAffected == 0 {
		return nil
	}

	return &user
}

func (r *UserRepository) FindByID(id uint) (*models.Users, error) {
	var user models.Users
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
