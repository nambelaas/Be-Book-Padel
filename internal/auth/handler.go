package auth

import (
	"Be-Book-Padel/database"
	"Be-Book-Padel/models"
	refreshtoken "Be-Book-Padel/models/repository/refresh_token"
	"Be-Book-Padel/models/repository/user"
	authServices "Be-Book-Padel/models/service/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	UserRepository         = user.UserRepository{}
	RefreshTokenRepository = refreshtoken.RefreshTokenRepository{}
	AuthService            = authServices.NewAuthService{
		UserRepo:         UserRepository,
		RefreshTokenRepo: RefreshTokenRepository,
	}
)

type AuthHandler struct {
	DB *gorm.DB
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{DB: database.DB}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		FirstName string `json:"first_name" binding:"required`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=6"`
		Gender    string `json:"gender" binding:"required"`
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	userModel := &models.Users{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Gender:    req.Gender,
	}

	err := AuthService.Register(userModel)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "user registered successfully",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	access, refresh, err := AuthService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	accessToken, refreshToken, err := AuthService.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *AuthHandler) RegisterStaff(c *gin.Context) {
	var req struct {
		FirstName string `json:"first_name" binding:"required`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=6"`
		Gender    string `json:"gender" binding:"required"`
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	userModel := &models.Users{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Gender:    req.Gender,
		Role:      models.Staff,
	}

	err := AuthService.Register(userModel)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "user registered successfully",
	})
}

func (h *AuthHandler) RegisterAdmin(c *gin.Context) {
	var req struct {
		FirstName string `json:"first_name" binding:"required`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=6"`
		Gender    string `json:"gender" binding:"required"`
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	userModel := &models.Users{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Gender:    req.Gender,
		Role:      models.Admin,
	}

	err := AuthService.Register(userModel)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "user registered successfully",
	})
}
