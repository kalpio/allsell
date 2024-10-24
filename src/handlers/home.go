package handlers

import (
	"github.com/kalpio/allsell/src/views/home"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HomeHandler struct{}

func (h HomeHandler) IndexGet(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.JSON(http.StatusOK, "dupa")
	}
	username := sess.Values["username"].(string)
	return render(c, home.IndexGet(username))
}
