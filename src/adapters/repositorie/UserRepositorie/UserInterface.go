package userrepositorie

import (
	"context"

	domain "github.com/moisesPompilio/projeto_x/src/internal/domain/User"
)

type UserRepositorie interface {
	Create(ctx context.Context, newUser domain.User) error
	GetUserByEmail(ctx context.Context, email string) ([]domain.User, error)
	// GetByID(ctx context.Context, ID int64) (*domain.UserWithoutPassword, error)
	// GetAll(ctx context.Context) ([]domain.UserWithoutPassword, error)
}
