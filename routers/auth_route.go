package routers

import (
	"hacktiv8-chapter-three-session-two/controllers"
	"hacktiv8-chapter-three-session-two/repository"
	"hacktiv8-chapter-three-session-two/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoute(router *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(*userService)

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.Registration)
		userRouter.POST("/login", userController.Login)
	}
}
