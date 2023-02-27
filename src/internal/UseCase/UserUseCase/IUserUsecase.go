package userusecase

import (
	"context"

	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/input"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/output"
)

type IUserUsecase interface {
	LoginUseCase(data input.Login, ctx context.Context) (output.Token, error)
	DeleteUsersById(data string, ctx context.Context) error
	// GetByID(ctx context.Context, ID int64) (*domain.UserWithoutPassword, error)
	// GetAll(ctx context.Context) ([]domain.UserWithoutPassword, error)
	CreateUserUseCase(create input.CreateUserDTO, ctx context.Context) error
}
