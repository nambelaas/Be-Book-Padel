package field

import (
	"Be-Book-Padel/models"
	"Be-Book-Padel/models/repository/field"
)

func AuthService(fieldRepo field.FieldRepository) FieldServiceInterface {
	return &NewFieldService{
		FieldRepo: fieldRepo,
	}
}

func (s *NewFieldService) Create(data *models.Field) (id uint, error error) {
	result, err := s.FieldRepo.Create(data)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (s *NewFieldService) Update(data *models.Field) (*models.Field, error) {
	updatedField, err := s.FieldRepo.Update(data)
	if err != nil {
		return nil, err
	}

	return updatedField, nil
}

func (s *NewFieldService) GetById(id uint) (*models.Field, error) {
	field, err := s.FieldRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return field, nil
}

func (s *NewFieldService) GetAllFields() ([]models.Field, error) {
	fields, err := s.FieldRepo.GetAllFields()
	if err != nil {
		return nil, err
	}
	return fields, nil
}

func (s *NewFieldService) Delete(id uint) error {
	err := s.FieldRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
