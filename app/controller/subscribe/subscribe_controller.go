package subscribe_controller

import "github.com/labstack/echo/v4"

type SubscribeController interface {
	SubscribeChannel(ctx echo.Context) error
	UnsubscribeChannel(ctx echo.Context) error
}
