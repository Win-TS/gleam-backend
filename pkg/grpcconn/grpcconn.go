package grpcconn

import (
	"errors"
	"log"
	"net"

	"github.com/Win-TS/gleam-backend.git/config"
	userPb "github.com/Win-TS/gleam-backend.git/modules/user/userPb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	GrpcClientFactoryHandler interface {
		User() userPb.UserGrpcServiceClient
	}

	grpcClientFactory struct {
		client *grpc.ClientConn
	}
)

func (g *grpcClientFactory) User() userPb.UserGrpcServiceClient {
	return userPb.NewUserGrpcServiceClient(g.client)
}

func NewGrpcClient(host string) (GrpcClientFactoryHandler, error) {
	opts := make([]grpc.DialOption, 0)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	clientConn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Printf("Error - Grpc client connection failed: %v", err.Error())
		return nil, errors.New("error - grpc client connection failed")
	}

	return &grpcClientFactory{client: clientConn}, nil
}

func NewGrpcServer(cfg *config.Config, host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)
	grpcServer := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Error - Failed to listen: %v", err)
	}

	return grpcServer, lis
}
