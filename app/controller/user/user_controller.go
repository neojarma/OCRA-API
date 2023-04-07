package user_controller

import "github.com/labstack/echo/v4"

type UserController interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
	Logout(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
}
