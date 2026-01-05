package field

import "Be-Book-Padel/models"

type FieldRepositoryInterface interface {
	Create(data *models.Field) (id uint, error error)
	Update(data *models.Field) (*models.Field, error)
	GetById(id uint) (*models.Field, error)
	GetAllFields() ([]models.Field, error)
	Delete(id uint) error
}

type FieldRepository struct{}
