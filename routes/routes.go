package routes

import (
	"Be-Book-Padel/internal/auth"
	"Be-Book-Padel/internal/user"
	"Be-Book-Padel/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	authHandler := auth.NewAuthHandler()

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/refresh-token", authHandler.RefreshToken)
		authRoutes.GET("/:id", authHandler.RefreshToken)
	}

	userHandler := user.NewUserHandler()

	userRoutes := r.Group("/api/user")
	{
		userRoutes.GET("/profile", middleware.CheckJwt(), userHandler.GetUserProfile)
	}
}
