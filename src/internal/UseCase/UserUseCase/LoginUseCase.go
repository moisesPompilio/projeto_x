package userusecase

import (
	"context"
	"errors"

	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/input"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/output"
)

func (usecase *Userusecase) LoginUseCase(data input.Login, ctx context.Context) (output.Token, error) {
	user, err := usecase.UserRepositorie.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return output.Token{}, err
	}
	if user.Email != data.Email || user.Password != data.Password {
		return output.Token{}, errors.New("E-mail ou senha inválidos.")
	}

	return output.Token{Token: "todo bem", RefreshToken: "ai sim"}, nil
}
