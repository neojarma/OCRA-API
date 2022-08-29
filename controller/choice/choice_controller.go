package choice_controller

import "github.com/labstack/echo/v4"

type ChoiceController interface {
	CreateLike(ctx echo.Context) error
	CreateDislike(ctx echo.Context) error
	DeleteLike(ctx echo.Context) error
	DeleteDislike(ctx echo.Context) error
}
