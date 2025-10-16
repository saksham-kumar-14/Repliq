package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func writeJSON(c echo.Context, statusCode int, data any) error {
	return c.JSON(statusCode, data)
}

func readJSON(c echo.Context, data any) error {
	maxBytes := int64(1_048_578) // 1 MB limit
	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, maxBytes)

	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(data); err != nil {
		return err
	}

	// optional: validate struct
	if err := Validate.Struct(data); err != nil {
		return err
	}

	return nil
}

func writeJSONError(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, echo.Map{
		"error": message,
	})
}
