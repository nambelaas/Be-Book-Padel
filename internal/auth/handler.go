package auth

import "github.com/gin-gonic/gin"

type Handler struct {
	Service *Service
}

func (h *Handler) Register(c *gin.Context) {
	var req struct {
		FirstName string `json:"first_name" binding:"required`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=6"`
		Gender    string `json:"gender" binding:"required"`
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
	}

	accessToken, refreshToken,err := h.Service.Register(req.FirstName, req.LastName, req.Email, req.Password, req.Gender)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"message": "user registered successfully",
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	access, refresh, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}
