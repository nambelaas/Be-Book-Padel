package auth

import (
	"errors"
	"time"

	"Be-Book-Padel/models"
	refreshtoken "Be-Book-Padel/models/repository/refresh_token"
	"Be-Book-Padel/models/repository/user"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UserRepo        user.UserRepository
	RefreshTokenRpo refreshtoken.RefreshTokenRepository
}

func (s *Service) Register(firstName, lastName, email, password, gender string) (string, string, error) {
	_, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	user := models.Users{
		Email:     email,
		Password:  string(hashedPassword),
		FirstName: firstName,
		LastName:  lastName,
		Gender:    gender,
	}

	err = s.UserRepo.Create(&user)
	if err != nil {
		return "", "", err
	}
	accessToken, _ := GenerateAccessToken(user.ID, user.Role)
	refreshToken := GenerateRandomToken()
	hashRefreshToken := HashToken(refreshToken)

	err = s.RefreshTokenRpo.Create(&models.RefreshToken{
		UserID:    user.ID,
		Token:     hashRefreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	})

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *Service) Login(email, password string) (string, string, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", "", err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", "", errors.New("Invalid credentials")
	}

	accessToken, _ := GenerateAccessToken(user.ID, user.Role)
	refreshToken := GenerateRandomToken()
	hashRefreshToken := HashToken(refreshToken)

	err = s.RefreshTokenRpo.Create(&models.RefreshToken{
		UserID:    user.ID,
		Token:     hashRefreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	})

	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
