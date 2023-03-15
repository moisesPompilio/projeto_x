package userusecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/output"
)

func (usecase *Userusecase) GetByID(ctx context.Context, ID uuid.UUID) (*output.UserOutpDTO, error) {
	user, err := usecase.UserRepositorie.GetByID(ctx, ID)
	if err != nil {
		return user, err
	}
	return user, nil
}
