package auth_middleware

import (
	"net/http"
	"ocra_server/model/response"
	cookie_service "ocra_server/service/cookie"
	session_service "ocra_server/service/session"
	"sync"

	"github.com/labstack/echo/v4"
)

type AuthMiddlewareImpl struct {
	SessionService session_service.SessionService
}

func NewAuthMiddleware(service session_service.SessionService) AuthMiddleware {
	var doOnce sync.Once
	middleware := new(AuthMiddlewareImpl)

	doOnce.Do(func() {
		middleware = &AuthMiddlewareImpl{
			SessionService: service,
		}
	})

	return middleware
}

func (middleware *AuthMiddlewareImpl) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cookie, err := c.Cookie("session-id")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, &response.EmptyObjectDataResponse{
				Status:  response.StatusFailed,
				Message: response.MessageMissingSessionId,
			})
		}

		sessionId := cookie.Value

		_, err = middleware.SessionService.CheckActiveSession(sessionId)
		if err != nil {
			cookieService := cookie_service.NewCookieService().DestroyCookie("session-id")
			c.SetCookie(cookieService)

			return c.JSON(http.StatusUnauthorized, &response.EmptyObjectDataResponse{
				Status:  response.StatusFailed,
				Message: response.MessageInvalidSession,
			})
		}

		return next(c)
	}
}
