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
		return user, err
	}
	conection := userrepositorie.NewUserRepository(db.GetReaderMySQL(), db.GetWriterMySQL())
	err = conection.Create(ctx, user)
	if err != nil {
		return user, err
	}
	return user, nil

}
