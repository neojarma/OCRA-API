package auth_middleware

import "github.com/labstack/echo/v4"

type AuthMiddleware interface {
	Auth(next echo.HandlerFunc) echo.HandlerFunc
}
