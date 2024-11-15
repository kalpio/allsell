package handlers

import (
	"context"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func CurrentUserName(ctx context.Context) string {
	value := ctx.Value("UserName")
	if value != nil {
		return value.(string)
	}

	return ""
}
