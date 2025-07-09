package auth

import (
	"os"
	"fmt"
	"time"
	"context"

	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"

	pb "auth-service/pb"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *Server) GenJWT(ctx context.Context, req *pb.JWTReq) (*pb.JWTRes, error) {
	userID := req.GetUserID()
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(24*time.Hour).Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	return &pb.JWTRes{Token: token}, nil
}

func (s *Server) HashPd(ctx context.Context, req *pb.HashReq) (*pb.HashRes, error) {
	pd := req.GetPassword()
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(pd), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &pb.HashRes {Hash: string(bytes)}, nil
}

func (s *Server) ComparePd(ctx context.Context, req *pb.CompareReq) (*pb.CompareRes, error) {
	pd := req.GetPassword()
	hash := req.GetHash()
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash), []byte(pd))
	if err != nil {
		return nil, err
	}
	return &pb.CompareRes {
		Compare: true,
	}, nil
}

func (s *Server) ExtUserID(ctx context.Context, req *pb.ExtReq) (*pb.ExtRes, error) {
	tokenStr := req.GetToken()

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing metod: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, fmt.Errorf("user_id not found or invalid type")
	}
	return &pb.ExtRes{UserID: userID}, nil
}
