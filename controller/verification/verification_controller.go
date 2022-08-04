package verification_controller

import "github.com/labstack/echo/v4"

type VerificationController interface {
	CreateVerificationToken(ctx echo.Context) error
	ValidateVerificationToken(ctx echo.Context) error
}
