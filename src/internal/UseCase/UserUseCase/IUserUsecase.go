package userusecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/input"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/output"
)

type IUserUsecase interface {
	LoginUseCase(data input.Login, ctx context.Context) (output.Token, error)
	DeleteUsersById(data string, ctx context.Context) error
	GetByID(context.Context, uuid.UUID) (*output.UserOutpDTO, error)
	CreateUserUseCase(create input.CreateUserDTO, ctx context.Context) error
}
