package comment_controller

import "github.com/labstack/echo/v4"

type CommentController interface {
	GetComment(ctx echo.Context) error
	CreateComment(ctx echo.Context) error
	UpdateComment(ctx echo.Context) error
	DeleteComment(ctx echo.Context) error
}
