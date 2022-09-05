package watchlater_controller

import (
	"github.com/labstack/echo/v4"
)

type WatchLaterController interface {
	GetAllWatchLaterRecords(ctx echo.Context) error
	CreateWatchLaterRecord(ctx echo.Context) error
	DeleteWatchLaterRecord(ctx echo.Context) error
}