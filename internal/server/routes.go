package server

import (
	"net/http"

	"github.com/MyongJi-StudentVote/StudentVote-Admin/internal/handler"
)

const BASE_URL = "http://13.125.244.156:8080/api/v1/%s"

func (app *Application) routes() {
	client := &http.Client{}
	signupHandler := handler.InitSignupHandler(client, BASE_URL)
	homeHandler := handler.InitHomeHandler(client, BASE_URL)
	app.Server.Static("/static", "static")
	app.Server.GET("/", homeHandler.Home)
	app.Server.POST("/signin", homeHandler.HandleSignin)
	app.Server.GET("/signup", signupHandler.Signup)
	app.Server.POST("/signup", signupHandler.HandleSignup)
	app.Server.GET("/main-home", handler.InitMainHomeHandler().MainHome)
}
