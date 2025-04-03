package api

import (
	"context"
	"log"

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
	log.Printf("Received GetUser request for ID: %s", req.UserId)

	user, err := s.Repo.GetUserByID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserResponse{
		UserId: user.ID,
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
		userList = append(userList, &userpb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return &userpb.ListUsersResponse{Users: userList}, nil
}

// Update user details
func (s *UserServer) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	log.Printf("Received UpdateUser request for ID: %s", req.UserId)

	updatedUser, err := s.Repo.UpdateUser(ctx, &domain.User{
		ID:    req.UserId,
		Name:  *req.Name,
		Email: *req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		UserId:    updatedUser.ID,
		Name:      updatedUser.Name,
		Email:     updatedUser.Email,
		UpdatedAt: updatedUser.UpdatedAt.String(),
	}, nil
}

// Delete user
func (s *UserServer) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	log.Printf("Received DeleteUser request for ID: %s", req.UserId)

	err := s.Repo.DeleteUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &userpb.DeleteUserResponse{
		Success: true,
	}, nil
}
