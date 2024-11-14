package server

import "github.com/MyongJi-StudentVote/StudentVote-Admin/internal/handler"

func (app *Application) routes() {
	app.Server.Static("/static", "static")
	app.Server.GET("/health_check", handler.HealthCheck)
}
