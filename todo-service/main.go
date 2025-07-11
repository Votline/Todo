package main

import (
	"os"
	"log"

	"google.golang.org/grpc"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/credentials/insecure"

	"todo-service/internal/db"
	"todo-service/internal/repo"
	"todo-service/internal/routers"
	"todo-service/internal/handlers"
	auth "todo/auth-service/pb"
)

func main() {
	conn, err := grpc.NewClient(
		os.Getenv("AUTH_HOST")+":"+os.Getenv("AUTH_PORT"),
		grpc.WithTransportCredentials(insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatalln("Error when trying to create conn with auth-service", err)
	}
	tdb := db.InitDB()
	authClient := auth.NewAuthServiceClient(conn)
	h := &handlers.Handler {
		AuthClient: authClient,
		Tdb: tdb,
		Trs: repo.NewTRS(tdb),
	}


	e := echo.New()
	routers.Setup(e, h)
	e.Logger.Fatal(e.Start(":"+os.Getenv("HTTP_PORT")))
}
