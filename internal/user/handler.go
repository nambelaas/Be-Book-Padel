package user

import (
	"Be-Book-Padel/database"
	"Be-Book-Padel/models/repository/user"
	userServices "Be-Book-Padel/models/service/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	UserRepo    = user.UserRepository{}
	UserService = userServices.NewUserService{
		UserRepo: UserRepo,
	}
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler() *UserHandler {
	return &UserHandler{DB: database.DB}
}

func (h *UserHandler) GetUserProfile(c *gin.Context){
	userId := c.GetUint("user_id")
	user, err := UserService.GetUserByID(userId)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}
