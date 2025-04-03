package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.70

import (
	"context"
	"log"

	"github.com/godfreyowidi/tiwabet-backend/gql-gateway/graph"
	"github.com/godfreyowidi/tiwabet-backend/gql-gateway/model"
	model1 "github.com/godfreyowidi/tiwabet-backend/gql-gateway/model/dao"
	"github.com/godfreyowidi/tiwabet-backend/proto/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model1.User, error) {
	// Initialize gRPC client
	if r.GrpcClient == nil || r.GrpcConn.GetState() == connectivity.Shutdown {
		var err error
		r.GrpcConn, err = grpc.Dial("tiwabet-backend:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Failed to create gRPC client: %v", err)
			return nil, err
		}
		r.GrpcClient = userpb.NewUserServiceClient(r.GrpcConn)
	}

	// Call CreateUser gRPC method
	res, err := r.GrpcClient.CreateUser(ctx, &userpb.CreateUserRequest{
		Name:  input.Name,
		Email: input.Email,
	})
	if err != nil {
		log.Printf("Error calling CreateUser: %v", err)
		return nil, err
	}

	// Map gRPC response to Graphql model
	return &model1.User{
		ID:    res.UserId,
		Name:  res.Name,
		Email: res.Email,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUser) (*model1.User, error) {
	// Initialize gRPC
	if r.GrpcClient == nil || r.GrpcConn.GetState() == connectivity.Shutdown {
		var err error
		r.GrpcConn, err = grpc.Dial("tiwabet-backend:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Failed to create gRPC client: %v", err)
			return nil, err
		}
		r.GrpcClient = userpb.NewUserServiceClient(r.GrpcConn)
	}

	// Call UpdateUser gRPC method
	res, err := r.GrpcClient.UpdateUser(ctx, &userpb.UpdateUserRequest{
		UserId: id,
		Name:   input.Name,
		Email:  input.Email,
	})
	if err != nil {
		log.Printf("Error calling UpdateUser: %v", err)
		return nil, err
	}

	// Map gRPC response to gql model
	return &model1.User{
		ID:    res.UserId,
		Name:  res.Name,
		Email: res.Email,
	}, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	// --- Initialize gRPC ---
	if r.GrpcClient == nil || r.GrpcConn.GetState() == connectivity.Shutdown {
		var err error
		r.GrpcConn, err = grpc.Dial("tiwabet-backend:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Failed to create gRPC client: %v", err)
			return false, err
		}
		r.GrpcClient = userpb.NewUserServiceClient(r.GrpcConn)
	}

	// DeleteUser gRPC call
	res, err := r.GrpcClient.DeleteUser(ctx, &userpb.DeleteUserRequest{
		UserId: id,
	})
	if err != nil {
		log.Printf("Error calling DeleteUser: %v", err)
		return false, err
	}
	return res.Success, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
