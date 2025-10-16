package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *application) internalServerError(c echo.Context, err error) error {
	app.logger.Errorf("internal server error: method=%s path=%s error=%v",
		c.Request().Method, c.Request().URL.Path, err)

	return writeJSONError(c, http.StatusInternalServerError, "internal server error ;-;")
}

func (app *application) badRequestError(c echo.Context, err error) error {
	app.logger.Warnf("bad request: method=%s path=%s error=%v",
		c.Request().Method, c.Request().URL.Path, err)

	return writeJSONError(c, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundError(c echo.Context, err error) error {
	app.logger.Errorf("not found: method=%s path=%s error=%v",
		c.Request().Method, c.Request().URL.Path, err)

	return writeJSONError(c, http.StatusNotFound, "not found")
}

func (app *application) conflict(c echo.Context, err error) error {
	app.logger.Warnf("conflict: method=%s path=%s error=%v",
		c.Request().Method, c.Request().URL.Path, err)

	return writeJSONError(c, http.StatusConflict, "conflict occurred")
}

func (app *application) rateLimitExceededResponse(c echo.Context, retry string) error {
	app.logger.Warnf("rate limit exceeded: method=%s path=%s retry_after=%s",
		c.Request().Method, c.Request().URL.Path, retry)

	return writeJSONError(c, http.StatusTooManyRequests, "rate limit exceeded. Try after: "+retry)
}
