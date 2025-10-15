package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *application) healthChecker(c echo.Context) error {
	data := map[string]string{
		"status":  "ok",
		"version": version,
	}

	return c.JSON(http.StatusOK, data)
}
