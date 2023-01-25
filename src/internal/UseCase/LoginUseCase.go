package usecase

import (
	"context"
	"errors"

	userrepositorie "github.com/moisesPompilio/projeto_x/src/adapters/repositorie/UserRepositorie"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces"
	db "github.com/moisesPompilio/projeto_x/src/pkg/config/DB"
)

func LoginUseCase(data interfaces.Login, ctx context.Context) (interfaces.Token, error) {
	conection := userrepositorie.NewUserRepository(db.GetReaderPostgreSQL(), db.GetWriterPostgreSQL())
	user, err := conection.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return interfaces.Token{}, err
	}
	if user.Email != data.Email || user.Password != data.Password {
		return interfaces.Token{}, errors.New("Usuario invalido")
	}

	return interfaces.Token{Token: "todo bem", RefreshToken: "ai sim"}, nil
}
