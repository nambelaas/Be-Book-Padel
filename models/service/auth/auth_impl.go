package auth

import (
	"Be-Book-Padel/helper"
	"Be-Book-Padel/models"
	refreshtoken "Be-Book-Padel/models/repository/refresh_token"
	"Be-Book-Padel/models/repository/user"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func AuthService(userRepo user.UserRepository, refreshTokenRepo refreshtoken.RefreshTokenRepository) AuthServiceInterface {
	return &NewAuthService{
		UserRepo:         userRepo,
		RefreshTokenRepo: refreshTokenRepo,
	}
}

func (s *NewAuthService) Register(user *models.Users) error {
	emailExist := s.UserRepo.FindByEmail(user.Email)
	if emailExist != nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user = &models.Users{
		Email:     user.Email,
		Password:  string(hashedPassword),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Gender:    user.Gender,
	}

	err = s.UserRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *NewAuthService) Login(email, password string) (string, string, error) {
	user := s.UserRepo.FindByEmail(email)
	if user == nil {
		return "", "", errors.New("User not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", "", errors.New("Invalid credentials")
	}

	accessToken, _ := helper.GenerateAccessToken(user.ID, user.Role)
	refreshToken := helper.GenerateRandomToken()
	hashRefreshToken := helper.HashToken(refreshToken)

	err := s.RefreshTokenRepo.CreateRefreshToken(&models.RefreshToken{
		UserID:    user.ID,
		Token:     hashRefreshToken,
		ExpiresAt: time.Now().Add(30 * time.Minute),
	})

	if err != nil {
		return "", "", err
	}
	return accessToken, hashRefreshToken, nil
}

func (s *NewAuthService) RefreshToken(token string) (string, string, error) {
	refreshTokenData, err := s.RefreshTokenRepo.FindByToken(token)
	if err != nil {
		return "", "", err
	}

	if refreshTokenData == nil {
		return "", "", errors.New("Invalid refresh token")
	}

	if refreshTokenData.ExpiresAt.Before(time.Now()) {
		return "", "", errors.New("Refresh token expired")
	}

	user, err := s.UserRepo.FindByID(refreshTokenData.UserID)
	if err != nil {
		return "", "", err
	}

	if user == nil {
		return "", "", errors.New("User not found")
	}

	accessToken, _ := helper.GenerateAccessToken(user.ID, user.Role)
	newRefreshToken := helper.GenerateRandomToken()
	hashNewRefreshToken := helper.HashToken(newRefreshToken)

	err = s.RefreshTokenRepo.CreateRefreshToken(&models.RefreshToken{
		UserID:    user.ID,
		Token:     hashNewRefreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	})

	if err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken, nil

}
