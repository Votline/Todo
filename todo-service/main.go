package main

import (
	"os"

	"github.com/labstack/echo/v4"

	"todo-service/internal/routers"
)

func main() {
	port := ":"
	port += os.Getenv("HTTP_PORT")
	
	e := echo.New()
	routers.Setup(e)
	e.Logger.Fatal(e.Start(port))
}
