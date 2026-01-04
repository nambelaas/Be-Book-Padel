package refreshtoken

import (
	"Be-Book-Padel/database"
	"Be-Book-Padel/models"
)

func NewRefreshTokenRepository() RefreshTokenRepositoryInterface{
	return &RefreshTokenRepository{}
}

func (r *RefreshTokenRepository) CreateRefreshToken(token *models.RefreshToken) error{
	result := database.DB.Create(token)

	return result.Error
}

func (r *RefreshTokenRepository) FindByToken(token string) (*models.RefreshToken, error){
	var refreshToken models.RefreshToken
	result := database.DB.Where("token = ?", token).First(&refreshToken)
	if result.Error != nil {
		return nil, result.Error
	}

	return &refreshToken, nil
}