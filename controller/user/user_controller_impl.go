package user_controller

import (
	"log"
	"net/http"
	"ocra_server/model/request"
	user_service "ocra_server/service/user"
	"sync"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	Service user_service.UserService
}

func NewUserController(service user_service.UserService) UserController {
	var doOnce sync.Once
	controller := new(UserControllerImpl)

	doOnce.Do(func() {
		controller = &UserControllerImpl{
			Service: service,
		}
	})

	return controller
}

func (controller *UserControllerImpl) Login(ctx echo.Context) error {
	req := new(request.UserRequest)
	ctx.Bind(req)

	if _, err := controller.Service.ValidateLogin(req); err != nil {
		return ctx.String(http.StatusBadRequest, "false")
	}
	return ctx.String(http.StatusOK, "ok")
}

func (controller *UserControllerImpl) Register(ctx echo.Context) error {
	req := new(request.UserRequest)
	ctx.Bind(req)
	if _, err := controller.Service.Register(req); err != nil {
		log.Println(err.Error())
		return ctx.String(http.StatusInternalServerError, "false")
	}

	return ctx.String(http.StatusOK, "ok")

}

func (controller *UserControllerImpl) UpdateUser(ctx echo.Context) error {
	req := new(request.UserRequest)
	ctx.Bind(req)
	if _, err := controller.Service.UpdateUser(req); err != nil {
		return ctx.String(http.StatusInternalServerError, "false")
	}

	return ctx.String(http.StatusOK, "ok")
}
