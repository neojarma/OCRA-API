package renew_session

import (
	"net/http"
	"ocra_server/model/response"
	session_service "ocra_server/service/session"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type RenewSessionImpl struct {
	SessionService session_service.SessionService
}

func NewRenewSession(service session_service.SessionService) RenewSession {
	var doOnce sync.Once
	middleware := new(RenewSessionImpl)

	doOnce.Do(func() {
		middleware = &RenewSessionImpl{
			SessionService: service,
		}
	})

	return middleware
}

func (middleware *RenewSessionImpl) RenewSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session-id")
		if err != nil {
			return next(c)
		}

		sessionId := cookie.Value
		if sessionId == "" {
			return next(c)
		}

		session, err := middleware.SessionService.CheckActiveSession(sessionId)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, &response.EmptyObjectDataResponse{
				Status:  response.StatusFailed,
				Message: response.MessageInvalidSession,
			})
		}

		c.Request().Header.Add("user-id", session.UserId)

		isSessionNearlyExpired := time.Now().UnixMilli()-session.ExpiresAt <= int64(time.Minute)*30
		if isSessionNearlyExpired {
			middleware.SessionService.UpdateExpiresSession(sessionId)
			return next(c)
		}

		return next(c)
	}
}
