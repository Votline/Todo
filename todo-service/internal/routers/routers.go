package routers

import (
	"os"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	mdwr "github.com/labstack/echo/v4/middleware"

	auth "todo/auth-service/pb"
	"todo-service/internal/handlers"
)

func Setup(e *echo.Echo, h *handlers.Handler) {
	origin := os.Getenv("CORS_ALLOW_ORIGINS")

	e.Use(mdwr.CORSWithConfig(mdwr.CORSConfig{
		AllowOrigins: []string{origin},
		AllowMethods: []string{
			"GET", "POST",
			"DELETE", "OPTIONS",
		},
	}))

	jwtMdwr := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "/api/todos/reg"},
		SuccessHandler: func(c echo.Context) {
			token, ok := c.Get("user").(*jwt.Token)
			if !ok {
				c.Set("auth_error", "Invalid or missing token")
				return
			}
			res, err := h.AuthClient.ExtUserID(
				c.Request().Context(),
				&auth.ExtReq{Token: token.Raw})
			if err != nil {
				c.Set("auth_error",
					"Auth service rejected token: "+err.Error())
				return
			}
			c.Set("userID", res.UserID)
		},
	})
	e.Use(jwtMdwr)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func (c echo.Context) error {
			if err, ok := c.Get("auth_error").(error); ok {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Auth error: " + err.Error())
			}
			return next(c)
		}
	})
	e.POST("api/todos/reg", h.AddUser)
	e.POST("api/todos/task", h.AddTask)
}
