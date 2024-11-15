package middleware

import (
	"context"
	"github.com/kalpio/allsell/src/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

type AuthorizationConfig struct {
	Skipper    middleware.Skipper
	SessionKey string
}

var DefaultAuthorizationConfig = AuthorizationConfig{
	Skipper:    middleware.DefaultSkipper,
	SessionKey: "UserName",
}

func Authorize() echo.MiddlewareFunc {
	return AuthorizeWithConfig(DefaultAuthorizationConfig)
}

func AuthorizeWithConfig(config AuthorizationConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultAuthorizationConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			url := req.URL
			path := url.Path

			if strings.HasPrefix(path, "/user") {
				return next(c)
			}

			var login = func(c echo.Context) error {
				return c.Redirect(http.StatusMovedPermanently, "/user/login")
			}
			value := session.Get[string](c, config.SessionKey)
			if isNone, _ := value.IsNone(); isNone {
				// log err
				return login(c)
			}
			if len(value.Unwrap()) == 0 {
				return login(c)
			}

			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), "UserName", value.Unwrap())))
			return next(c)
		}
	}
}
