package utils

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(component templ.Component, ctx echo.Context, code int) error {
	ctx.Response().Status = code
	return component.Render(ctx.Request().Context(), ctx.Response())
}
