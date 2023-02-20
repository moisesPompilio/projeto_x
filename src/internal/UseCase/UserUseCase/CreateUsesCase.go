package userusecase

import (
	"context"

	"github.com/moisesPompilio/projeto_x/src/internal/domain"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/input"
)

func (usecase *Userusecase) GetUserUseCase(create input.CreateUserDTO, ctx context.Context) error {
	user, err := domain.NewUser(create)
	if err != nil {
		return err
	}

	err = usecase.UserRepositorie.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil

}
