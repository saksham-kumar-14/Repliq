package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/saksham-kumar-14/Repliq/backend/internal/store"
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
	store  store.Storage
}

func (app *application) mount() *echo.Echo {
	e := echo.New()

	e.GET("/v1/health", app.healthChecker)

	users := e.Group("/v1/user")
	users.POST("", app.registerUserHandler)
	users.GET("/:id", app.getUserHandler)
	users.POST("/login", app.loginUserHandler)

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
