package main

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/saksham-kumar-14/Repliq/backend/internal/auth"
	"github.com/saksham-kumar-14/Repliq/backend/internal/env"
)

var JwtSecret = []byte(env.GetString("JWT_SECRET", "secret"))

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token format"})
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "unexpected signing method")
			}
			return JwtSecret, nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token claims"})
		}

		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Set("email", claims["email"].(string))

		return next(c)
	}
}

func TokenApi(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := ""
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	}

	if tokenString == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "token missing"})
	}

	token, err := auth.ParseJWT(tokenString)
	if err != nil || !token.Valid {
		return c.JSON(http.StatusOK, echo.Map{
			"valid": false,
			"user":  nil,
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusOK, echo.Map{
			"valid": false,
			"user":  nil,
		})
	}

	user := map[string]interface{}{
		"user_id": claims["user_id"],
		"email":   claims["email"],
	}

	return c.JSON(http.StatusOK, echo.Map{
		"valid": true,
		"user":  user,
	})
}
