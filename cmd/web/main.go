package main

import (
	"context"
	"io"

	"github.com/MyongJi-StudentVote/StudentVote-Admin/templates"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Render(component templ.Component, ctx context.Context, w io.Writer) error {
	return component.Render(ctx, w)
}

func healthCheck(ctx echo.Context) error {
	home := templates.Home()
	return Render(home, context.Background(), ctx.Response())
}

func main() {
	srv := echo.New()
	srv.Static("/static", "static")
	srv.GET("/health_check", healthCheck)
	srv.Logger.SetLevel(log.INFO)
	if err := srv.Start(":8080"); err != nil {
		log.Fatalf("Error occured while starting server : %v", err)
	}
}
