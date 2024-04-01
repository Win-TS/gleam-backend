package middlewareRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Win-TS/gleam-backend.git/pkg/grpcconn"
	authPb "github.com/Win-TS/gleam-backend.git/modules/auth/authPb"
)

type (
	MiddlewareRepositoryService interface{
		VerifyToken(pctx context.Context, grpcUrl, token string) error
	}

	middlewareRepository        struct{}
)

func NewMiddlewareRepository() MiddlewareRepositoryService {
	return &middlewareRepository{}
}

func (r *middlewareRepository) VerifyToken(pctx context.Context, grpcUrl, token string) error {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	conn, err := grpcconn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("error - gRPC connection failed: %s", err.Error())
		return errors.New("error: gRPC connection failed")
	}

	result, err := conn.Auth().VerifyToken(ctx, &authPb.VerifyTokenReq{Token: token})
	if err != nil {
		log.Printf("error - VerifyToken failed: %s", err.Error())
		return errors.New("error: access token invalid")
	}

	if !result.Success {
		return errors.New("error: access token invalid")
	}

	return nil
}