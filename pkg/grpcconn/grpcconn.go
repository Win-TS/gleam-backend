package grpcconn

import (
	"errors"
	"log"
	"net"

	// inventoryPb "github.com/Win-TS/go-course-microservice-shop-tutorial/modules/inventory/inventoryPb"
	// itemPb "github.com/Win-TS/go-course-microservice-shop-tutorial/modules/item/itemPb"
	// playerPb "github.com/Win-TS/go-course-microservice-shop-tutorial/modules/player/playerPb"
	"github.com/Win-TS/gleam-backend.git/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	GrpcClientFactoryHandler interface {
		//User() userPb.UserGrpcServiceClient
	}

	grpcClientFactory struct {
		client *grpc.ClientConn
	}

	grpcAuth struct {
	}
)

// func (g *grpcClientFactory) User() userPb.UserGrpcServiceClient {
// 	return userPb.NewUserGrpcServiceClient(g.client)
// }

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
