package channel_controller

import "github.com/labstack/echo/v4"

type ChannelController interface {
	CreateChannel(ctx echo.Context) error
	DetailChannel(ctx echo.Context) error
}
