package renew_session

import "github.com/labstack/echo/v4"

type RenewSession interface {
	RenewSession(next echo.HandlerFunc) echo.HandlerFunc
}
