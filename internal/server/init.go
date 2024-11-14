package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type Application struct {
	Server *echo.Echo
}

func NewApplication() *Application {
	e := echo.New()
	return &Application{
		Server: e,
	}
}

func gracefulShutdown(ctx context.Context, srv *echo.Echo, done chan bool) {
	signalCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	<-signalCtx.Done()

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	if err := srv.Shutdown(timeoutCtx); err != nil {
		log.Fatalf("Fail to shutdown server in given time: %v", err)
	}
	log.Print("Server shutting down... \n")
	done <- true
}

func (app *Application) Run() {
	done := make(chan bool, 1)
	app.routes()
	go func() {
		if err := app.Server.Start(":8080"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error occured while starting server : %v", err)
		}
	}()
	gracefulShutdown(context.Background(), app.Server, done)
	<-done
	log.Print("Server shutdown properly.\n")
}
