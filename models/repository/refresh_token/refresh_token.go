package refreshtoken

import "Be-Book-Padel/models"

type RefreshTokenRepositoryInterface interface {
	CreateRefreshToken(data *models.RefreshToken) error
	FindByToken(token string) (*models.RefreshToken, error)
}

type RefreshTokenRepository struct{}
