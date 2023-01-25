package usecase

import (
	"context"

	userrepositorie "github.com/moisesPompilio/projeto_x/src/adapters/repositorie/UserRepositorie"
	"github.com/moisesPompilio/projeto_x/src/internal/domain"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces"
	db "github.com/moisesPompilio/projeto_x/src/pkg/config/DB"
)

func GetUserUseCase(create interfaces.CreateUserDTO, ctx context.Context) error {
	user, err := domain.NewUser(create)
	if err != nil {
		return err
	}
	conection := userrepositorie.NewUserRepository(db.GetReaderPostgreSQL(), db.GetWriterPostgreSQL())
	err = conection.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil

}
