package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type config struct {
	addr        string
	db          dbconfig
	env         string
	frontendURL string
}

type dbconfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type application struct {
	config config
	logger *zap.SugaredLogger
}

func (app *application) mount() *echo.Echo {
	e := echo.New()

	// Routes
	v1 := e.Group("/v1")
	{
		v1.GET("/health", app.healthChecker)
	}

	return e
}

func (app *application) run(e *echo.Echo) error {

	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      e,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	app.logger.Infow("The server is running", "addr", app.config.addr)
	return server.ListenAndServe()
}
