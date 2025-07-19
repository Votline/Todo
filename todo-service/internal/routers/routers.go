package routers

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-jwt/v4"
	mdwr "github.com/labstack/echo/v4/middleware"

	"todo-service/internal/handlers"
	"todo-service/internal/middleware"
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

	authMw := middleware.AuthMiddleware{
		AuthClient: h.AuthClient,
	}
	jwtMdwr := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "/api/todos/reg"},
		SuccessHandler: authMw.ExtUserID,
	})
	e.Use(jwtMdwr)
	e.Use(authMw.HandleError)
	e.POST("api/todos/reg", h.AddUser)
	
	e.GET("api/todos/task", h.GetTask)
	e.POST("api/todos/task", h.AddOrUpdTask)
	e.DELETE("api/todos/task", h.DelTask)
}
