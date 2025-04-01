package api

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/godfreyowidi/tiwabet-backend/domain"
	"github.com/godfreyowidi/tiwabet-backend/infra"
	"github.com/godfreyowidi/tiwabet-backend/infra/db"
	"github.com/godfreyowidi/tiwabet-backend/proto/userpb"
)

// Implements the UserService gRPC interface
type UserServer struct {
	Repo *infra.UserRepository
	userpb.UnimplementedUserServiceServer
	DB *db.PostgresDB
}

func NewUserService(userRepo *infra.UserRepository) *UserServer {
	return &UserServer{Repo: userRepo}
}

// fetch user by ID
func (s *UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	log.Printf("Received GetUser request for ID: %d", req.UserId)

	userIDStr := strconv.FormatInt(req.UserId, 10)

	user, err := s.Repo.GetUserByID(ctx, userIDStr)
	if err != nil {
		return nil, err
	}

	userIDInt, err := strconv.ParseInt(user.ID, 10, 64)
	if err != nil {
		log.Printf("Error converting user ID to int64: %v", err)
		return nil, err
	}
	return &userpb.GetUserResponse{
		UserId: userIDInt,
		Name:   user.Name,
		Email:  user.Email,
	}, nil
}

// handles user creation
func (s *UserServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	log.Printf("Received CreateUser request for Name: %s, Email: %s", req.Name, req.Email)

	user := &domain.User{
		Name:  req.Name,
		Email: req.Email,
	}
	createdUser, err := s.Repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{
		UserId:    createdUser.ID,
		Name:      createdUser.Name,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt.String(),
		UpdatedAt: createdUser.UpdatedAt.String(),
	}, nil
}

func (s *UserServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := s.Repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var userList []*userpb.User
	for _, user := range users {
		id, err := strconv.ParseInt(user.ID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert user ID %s to int64: %v", user.ID, err)
		}
		userList = append(userList, &userpb.User{
			Id:    id,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return &userpb.ListUsersResponse{Users: userList}, nil
}
