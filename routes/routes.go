package routes

import (
	"Be-Book-Padel/internal/auth"
	"Be-Book-Padel/internal/field"
	"Be-Book-Padel/internal/user"
	"Be-Book-Padel/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	authHandler := auth.NewAuthHandler()
	userHandler := user.NewUserHandler()
	fieldHandler := field.NewFieldHandler()

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/refresh-token", authHandler.RefreshToken)
		authRoutes.POST("/register-staff", middleware.CheckJwt(), middleware.AdminOrStaffOnly(), authHandler.RegisterStaff)
		authRoutes.POST("/register-admin", middleware.CheckJwt(), middleware.AdminOnly(), authHandler.RegisterAdmin)
	}

	userRoutes := r.Group("/api/user")
	{
		userRoutes.GET("/profile", middleware.CheckJwt(), userHandler.GetUserProfile)
	}

	fieldRoutes := r.Group("/api/field")
	{
		fieldRoutes.POST("/new", middleware.CheckJwt(), middleware.AdminOrStaffOnly(), fieldHandler.Create)
	}
}
