package handlers

import (
	"github.com/kalpio/allsell/src/views/home"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) IndexGet(c echo.Context) error {
	return render(c, home.IndexGet())
}
