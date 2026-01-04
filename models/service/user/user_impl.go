package user

import (
	"Be-Book-Padel/models"
	"Be-Book-Padel/models/repository/user"
)

func UserService(userRepo user.UserRepository) UserServiceInterface {
	return &NewUserService{
		UserRepo: userRepo,
	}
}

func (s *NewUserService) GetUserByID(id uint) (*models.Users, error) {
	user, err := s.UserRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
