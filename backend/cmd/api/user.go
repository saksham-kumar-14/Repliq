package main

import (
	"errors"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/saksham-kumar-14/Repliq/backend/internal/auth"
	"github.com/saksham-kumar-14/Repliq/backend/internal/store"
	"golang.org/x/crypto/bcrypt"
)

type registerUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type registerUserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	Avatar    string `json:"avatar"`
}

func (app *application) getUserHandler(c echo.Context) error {
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return app.internalServerError(c, err)
	}

	user, err := app.store.User.GetByID(c.Request().Context(), uint(userID))
	if err != nil {
		switch err {
		case store.ErrNotFound:
			return app.notFoundError(c, err)
		default:
			return app.internalServerError(c, err)
		}
	}

	return writeJSON(c, http.StatusOK, user)
}

func (app *application) registerUserHandler(c echo.Context) error {
	var req registerUserRequest
	if err := c.Bind(&req); err != nil {
		return app.badRequestError(c, err)
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		return app.badRequestError(c, errors.New("fields can't be empty"))
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return app.internalServerError(c, errors.New("failed to hash password"))
	}

	// random avatar
	avatarId := rand.IntN(100) + 1
	avatar_link := "https://avatar.iran.liara.run/public/" + strconv.Itoa(avatarId)
	user := &store.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashed,
		Avatar:   avatar_link,
	}

	err = app.store.User.Create(c.Request().Context(), user)
	if err != nil {
		return app.conflict(c, err)
	}

	resp := registerUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		Avatar:    user.Avatar,
	}

	return writeJSON(c, http.StatusCreated, resp)
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (app *application) loginUserHandler(c echo.Context) error {
	var req loginRequest
	if err := c.Bind(&req); err != nil {
		return app.badRequestError(c, err)
	}

	if req.Email == "" || req.Password == "" {
		return app.badRequestError(c, errors.New("fields can't be empty"))
	}

	user, err := app.store.User.VerifyUser(c.Request().Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return app.badRequestError(c, errors.New("invalid email or password"))
		}
		return app.internalServerError(c, err)
	}

	token, err := auth.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return app.internalServerError(c, errors.New("failed to generate token"))
	}

	resp := loginResponse{
		Token: token,
	}

	return writeJSON(c, http.StatusOK, resp)
}
