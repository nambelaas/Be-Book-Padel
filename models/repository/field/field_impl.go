package field

import (
	"Be-Book-Padel/database"
	"Be-Book-Padel/models"

	"gorm.io/gorm"
)

func NewFieldRepository() FieldRepositoryInterface {
	return &FieldRepository{}
}

func (r *FieldRepository) Create(data *models.Field) (id uint, error error) {
	result := database.DB.Create(data)
	if result.Error != nil {
		return 0, result.Error
	}

	database.DB.Save(data)

	return data.ID, nil
}

func (r *FieldRepository) Update(data *models.Field) (*models.Field, error) {
	result := database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *FieldRepository) GetById(id uint) (*models.Field, error) {
	var field models.Field
	result := database.DB.Preload("FieldPricings").First(&field, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &field, nil
}

func (r *FieldRepository) GetAllFields() ([]models.Field, error) {
	var fields []models.Field

	result := database.DB.Preload("FieldPricings").Find(&fields)
	if result.Error != nil {
		return nil, result.Error
	}

	return fields, nil
}

func (r *FieldRepository) Delete(id uint) error {
	result := database.DB.Select("FieldPricings").Delete(&models.Field{}, id)
	return result.Error
}
