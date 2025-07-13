package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"
	
	auth "todo/auth-service/pb"
)

type AuthMiddleware struct {
	AuthClient auth.AuthServiceClient
}

func (a *AuthMiddleware) ExtUserID (c echo.Context) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		c.Set("auth_error", "Invalid or missing token")
		return
	}
	res, err := a.AuthClient.ExtUserID(
		c.Request().Context(),
		&auth.ExtReq{Token: token.Raw})
	if err != nil {
		c.Set("auth_error",
			"Auth service rejected token:\n " + err.Error())
		return
	}
	c.Set("userID", res.UserID)
}

func (a *AuthMiddleware) HandleError(next echo.HandlerFunc) echo.HandlerFunc {
	return func (c echo.Context) error {
		if msg, ok := c.Get("auth_error").(string); ok {
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"Auth error: " + msg)
		}
		return next(c)
	}
}
