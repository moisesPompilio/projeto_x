package usecase

import (
	"context"

	userrepositorie "github.com/moisesPompilio/projeto_x/src/adapters/repositorie/UserRepositorie"
	domain "github.com/moisesPompilio/projeto_x/src/internal/domain/User"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces"
	db "github.com/moisesPompilio/projeto_x/src/pkg/config/DB"
)

func UserUseCase(create interfaces.CreateUserDTO, ctx context.Context) (domain.User, error) {
	user, err := domain.NewUser(create)
	if err != nil {
		return domain.User{}, err
	}
	conection := userrepositorie.NewUserRepository(db.GetReaderPostgreSQL(), db.GetWriterPostgreSQL())
	err = conection.Create(ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil

}
