package userrepositorie

import (
	"context"

	"github.com/google/uuid"

	"github.com/moisesPompilio/projeto_x/src/internal/domain"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/output"
)

type IUserRepositorie interface {
	Create(ctx context.Context, newUser domain.User) error
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*output.UserOutpDTO, error)
	DeleteByID(ctx context.Context, ID string) error
}
