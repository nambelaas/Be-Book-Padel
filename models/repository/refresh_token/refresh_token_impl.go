package refreshtoken

import (
	"Be-Book-Padel/internal/database"
	"Be-Book-Padel/models"
)

func NewRefreshTokenRepository() RefreshTokenRepositoryInterface{
	return &RefreshTokenRepository{}
}

func (r *RefreshTokenRepository) Create(token *models.RefreshToken) error{
	result := database.DB.Create(token)

	return result.Error
}