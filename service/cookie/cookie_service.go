package cookie_service

import (
	"net/http"
	"time"
)

type CookieService interface {
	CreateCookie(key, value string, duration time.Time) *http.Cookie
	DestroyCookie(key string) *http.Cookie
}
