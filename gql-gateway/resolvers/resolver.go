package resolvers

import (
	"log"

	"github.com/godfreyowidi/tiwabet-backend/proto/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	GrpcConn   *grpc.ClientConn
	GrpcClient userpb.UserServiceClient
}

func NewResolver(grpcAddr string) (*Resolver, error) {
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
		return nil, err
	}

	client := userpb.NewUserServiceClient(conn)
	return &Resolver{
		GrpcConn:   conn,
		GrpcClient: client,
	}, nil
}

func (r *Resolver) Close() {
	if r.GrpcConn != nil {
		r.GrpcConn.Close()
	}
}
