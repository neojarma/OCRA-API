package watchlater_controller

import (
	"github.com/labstack/echo/v4"
)

type HistoryController interface {
	GetAllHistoryRecords(ctx echo.Context) error
	CreateHistoryRecord(ctx echo.Context) error
	DeleteHistoryRecord(ctx echo.Context) error
}
