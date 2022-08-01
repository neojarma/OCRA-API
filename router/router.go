package router

import (
	user_controller "ocra_server/controller/user"
	user_repository "ocra_server/repository/user"
	user_service "ocra_server/service/user"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Router(group *echo.Group, db *gorm.DB) {
	userRoute(group, db)
}

func userRoute(group *echo.Group, db *gorm.DB) {
	userRepo := user_repository.NewUserRepository(db)
	service := user_service.NewUserService(userRepo)
	controller := user_controller.NewUserController(service)

	group.POST("/auth/login", controller.Login)
	group.POST("/register", controller.Register)
	group.PATCH("/user/:id", controller.UpdateUser)
}
