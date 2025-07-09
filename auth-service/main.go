package main

import (
	"os"
	"log"
	"net"

	pb "auth-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	
	"auth-service/auth"
)

func main() {
	port := ":"+os.Getenv("AUTH_PORT")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("\n\nFailed to listen on port %s: %v\n\n",
			port, err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &auth.Server{})
	reflection.Register(s)

	s.Serve(lis)
}
