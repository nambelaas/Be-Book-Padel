package field

import (
	"Be-Book-Padel/models"
	"Be-Book-Padel/models/repository/field"
)

type FieldServiceInterface interface {
	Create(data *models.Field) (id uint, error error)
	Update(data *models.Field) (*models.Field, error)
	GetById(id uint) (*models.Field, error)
	GetAllFields() ([]models.Field, error)
	Delete(id uint) error
}

type NewFieldService struct {
	FieldRepo field.FieldRepository
}
