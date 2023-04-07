package cookie_service

import (
	"net/http"
	"sync"
	"time"
)

type cookieServiceImpl struct{}

func NewCookieService() CookieService {
	var doOnce sync.Once
	service := new(cookieServiceImpl)

	doOnce.Do(func() {
		service = &cookieServiceImpl{}
	})

	return service
}

func (service *cookieServiceImpl) CreateCookie(key, value string, duration time.Time) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = value
	cookie.Expires = duration
	cookie.HttpOnly = true
	cookie.Path = "/"
	// cookie.Domain = os.Getenv("CORS_DOMAIN")
	// cookie.SameSite = http.SameSiteLaxMode
	// cookie.Secure = true

	return cookie
}

func (service *cookieServiceImpl) DestroyCookie(key string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = ""
	cookie.HttpOnly = true
	cookie.Expires = time.Unix(0, 0)
	cookie.MaxAge = -1
	cookie.Path = "/"
	// cookie.Domain = os.Getenv("CORS_DOMAIN")
	// cookie.SameSite = http.SameSiteLaxMode
	// cookie.Secure = true

	return cookie
}
