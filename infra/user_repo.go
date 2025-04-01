package infra

import (
	"context"
	"log"

	"github.com/godfreyowidi/tiwabet-backend/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

// Initializes the repository
func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (*domain.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1`

	user := &domain.User{}
	err := r.DB.QueryRow(ctx, query, userID).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return nil, err
	}
	return user, nil
}

// Inserts a new user
func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `INSERT INTO users (name, email, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id, created_at, updated_at`
	err := r.DB.QueryRow(ctx, query, user.Name, user.Email).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) ListUsers(ctx context.Context) ([]*domain.User, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Printf("Error scanning user rows: %v", err)
			continue
		}
		users = append(users, user)
	}
	return users, nil
}
