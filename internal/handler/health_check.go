package handler

import (
	"net/http"

	"github.com/MyongJi-StudentVote/StudentVote-Admin/internal/utils"
	"github.com/MyongJi-StudentVote/StudentVote-Admin/templates"
	"github.com/labstack/echo/v4"
)

func HealthCheck(ctx echo.Context) error {
	home := templates.Home()
	return utils.Render(home, ctx, http.StatusOK)
}
