package routers

import (
	"os"

	"github.com/labstack/echo/v4"
	mdwr "github.com/labstack/echo/v4/middleware"

	"todo-service/internal/handlers"
)

func Setup(e *echo) {
	origin := os.Getenv("CORS_ALLOW_ORIGINS")

	e.Use(mdwr.CORSWithConfig(mdwr.CORSConfig{
		AllowOrigins: []string{origin},
		AllowMethods: []string{
			"GET", "POST",
			"DELETE", "OPTIONS",
		},
	}))

	open := e.Group("api/todos")
	open.GET("/", handlers.Hello)
}
