package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func param(c echo.Context, name string, def string) string {
	value := c.QueryParam(name)
	if len(value) == 0 {
		return def
	}

	return value
}
