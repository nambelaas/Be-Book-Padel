package fieldpricing

import "Be-Book-Padel/models"

type FieldPricingRepositoryInterface interface {
	Create(data *models.FieldPricing) (id uint, error error)
	Update(data *models.FieldPricing) (*models.FieldPricing, error)
	GetById(id uint) (*models.FieldPricing, error)
	GetAllFields() ([]models.FieldPricing, error)
	Delete(id uint) error
}

type FieldPricingRepository struct{}
