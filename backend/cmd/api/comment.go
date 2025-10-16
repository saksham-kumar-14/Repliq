package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/saksham-kumar-14/Repliq/backend/internal/store"
)

func (app *application) getAllCommentsHandler(c echo.Context) error {
	ctx := c.Request().Context()

	comments, err := app.store.Comment.GetAll(ctx)
	if err != nil {
		return app.internalServerError(c, err)
	}

	return writeJSON(c, http.StatusOK, comments)
}

func (app *application) getCommentHandler(c echo.Context) error {
	idParam := c.Param("id")
	commentID, err := strconv.Atoi(idParam)
	if err != nil {
		return app.internalServerError(c, err)
	}

	comment, err := app.store.User.GetByID(c.Request().Context(), uint(commentID))
	if err != nil {
		switch err {
		case store.ErrNotFound:
			return app.notFoundError(c, err)
		default:
			return app.internalServerError(c, err)
		}
	}

	return writeJSON(c, http.StatusOK, comment)
}

type createCommentRequest struct {
	ParentId int    `json:"parent_id"`
	Text     string `json:"text" validate:"required"`
	UserId   uint   `json:"user_id"`
}

func (app *application) createCommentHandler(c echo.Context) error {
	var req createCommentRequest
	if err := c.Bind(&req); err != nil {
		return app.badRequestError(c, err)
	}

	if req.Text == "" {
		return app.badRequestError(c, errors.New("text cannot be empty"))
	}

	comment := &store.Comment{
		ParentId:  req.ParentId,
		Text:      req.Text,
		UserId:    req.UserId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := app.store.Comment.Create(c.Request().Context(), comment)
	if err != nil {
		return app.internalServerError(c, err)
	}

	return writeJSON(c, http.StatusCreated, comment)
}

func (app *application) patchCommentHandler(c echo.Context) error {
	idParam := c.Param("id")
	commentID, err := strconv.Atoi(idParam)
	if err != nil {
		return app.badRequestError(c, err)
	}

	updates := make(map[string]interface{})
	if err := c.Bind(&updates); err != nil {
		return app.badRequestError(c, err)
	}

	updates["updated_at"] = time.Now()

	comment, err := app.store.Comment.PatchByID(c.Request().Context(), uint(commentID), updates)
	if err != nil {
		if err == store.ErrNotFound {
			return app.notFoundError(c, err)
		}
		return app.internalServerError(c, err)
	}

	return writeJSON(c, http.StatusOK, comment)
}

func (app *application) deleteCommentHandler(c echo.Context) error {
	idParam := c.Param("id")
	commentID, err := strconv.Atoi(idParam)
	if err != nil {
		return app.badRequestError(c, err)
	}

	err = app.store.Comment.DeleteByID(c.Request().Context(), uint(commentID))
	if err != nil {
		if err == store.ErrNotFound {
			return app.notFoundError(c, err)
		}
		return app.internalServerError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}
