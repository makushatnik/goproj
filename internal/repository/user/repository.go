package user

import (
	"context"
	"fmt"
	"goproj/internal/db/pg"
	"goproj/internal/entity"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(ctx context.Context, name string, age int64, verify bool) error
	List(ctx context.Context) ([]entity.User, error)
	GetUserByName(ctx context.Context, name string) (entity.User, error)
}

type repository struct {
	poolConn *pgxpool.Pool
}

func New(poolConn pg.Repository) UserRepository {
	return &repository{poolConn: poolConn.GetPoolConnection()}
}

func (r *repository) CreateUser(ctx context.Context, name string, age int64, verify bool) error {
	_, err := r.poolConn.Exec(ctx,
		"INSERT INTO users (created_at, name, age, verify) VALUES ($1, $2, $3, 4);",
		time.Now(), name, age, verify,
	)
	if err != nil {
		return fmt.Errorf("failed to exec CreateUser query: %w", err)
	}

	return nil
}

func (r *repository) List(ctx context.Context) ([]entity.User, error) {
	rows, err := r.poolConn.Query(ctx, "SELECT name, age, verify FROM users;")
	if err != nil {
		return nil, fmt.Errorf("failed to exec List query: %w", err)
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.Name, &user.Age, &user.Verify); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate rows: %w", err)
	}
	return users, nil
}

func (r *repository) GetUserByName(ctx context.Context, name string) (entity.User, error) {
	row := r.poolConn.QueryRow(ctx, "SELECT name, age, verify FROM users WHERE name = $1;", name)

	var user entity.User
	if err := row.Scan(&user.Name, &user.Age, &user.Verify); err != nil {
		return user, fmt.Errorf("failed to scan row: %w", err)
	}

	return user, nil
}
