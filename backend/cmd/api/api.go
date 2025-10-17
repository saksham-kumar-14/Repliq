package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	ratelimiter "github.com/saksham-kumar-14/Repliq/backend/internal/rateLimiter"
	"github.com/saksham-kumar-14/Repliq/backend/internal/store"
	"go.uber.org/zap"
)

type config struct {
	addr        string
	db          dbconfig
	env         string
	frontendURL string
	ratelimiter ratelimiterConfig
}

type dbconfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type ratelimiterConfig struct {
	ReqPerTimeFrame int
	Burst           int
	TimeFrame       time.Duration
}

type application struct {
	config      config
	logger      *zap.SugaredLogger
	store       store.Storage
	rateLimiter *ratelimiter.RateLimiter
}

func (app *application) mount() *echo.Echo {
	e := echo.New()

	// e.Use(app.rateLimiter.Limit)

	e.GET("/v1/api/token", TokenApi)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	health := e.Group("/v1/health")
	health.GET("", app.healthChecker)

	users := e.Group("/v1/user")
	users.POST("", app.registerUserHandler)
	users.GET("/:id", app.getUserHandler)
	users.POST("/login", app.loginUserHandler)

	posts := e.Group("/v1/post")
	posts.Use(JWTAuth)
	posts.GET("/", app.getAllCommentsHandler)
	posts.GET("/:id", app.getCommentHandler)
	posts.POST("/", app.createCommentHandler)
	posts.PATCH("/:id", app.patchCommentHandler)
	posts.DELETE("/:id", app.deleteCommentHandler)

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
