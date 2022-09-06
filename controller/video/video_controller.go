package video_controller

import (
	"github.com/labstack/echo/v4"
)

type VideoController interface {
	GetAllVideos(ctx echo.Context) error
	GetDetailVideos(ctx echo.Context) error
	CreateVideo(ctx echo.Context) error
	UpdateVideo(ctx echo.Context) error
	IncrementViews(ctx echo.Context) error
}
