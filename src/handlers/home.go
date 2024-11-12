package handlers

import (
	"github.com/kalpio/allsell/src/middleware"
	"github.com/kalpio/allsell/src/session"
	"github.com/kalpio/allsell/src/views/home"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) IndexGet(c echo.Context) error {
	value := session.Get[string](c, middleware.DefaultAuthorizationConfig.SessionKey)
	userName := value.UnwrapOr("Unknown")

	return render(c, home.IndexGet(userName))
}
