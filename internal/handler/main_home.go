package handler

import (
	"net/http"

	"github.com/MyongJi-StudentVote/StudentVote-Admin/internal/utils"
	"github.com/MyongJi-StudentVote/StudentVote-Admin/templates"
	"github.com/labstack/echo/v4"
)

type MainHomeHandler struct {
	Title string
}

func (m *MainHomeHandler) MainHome(ctx echo.Context) error {
	mainHome := templates.MainHome("메인 홈", "test@naver.com", "ICT 융합대학")
	return utils.Render(mainHome, ctx, http.StatusOK)
}

func InitMainHomeHandler() *MainHomeHandler {
	return &MainHomeHandler{}
}
