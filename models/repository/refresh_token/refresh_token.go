package refreshtoken

import "Be-Book-Padel/models"

type RefreshTokenRepositoryInterface interface {
	Create(data *models.RefreshToken) error
}

type RefreshTokenRepository struct{}
