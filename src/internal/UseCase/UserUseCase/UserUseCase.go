package userusecase

import (
	userrepositorie "github.com/moisesPompilio/projeto_x/src/adapters/repositorie/UserRepositorie"
	db "github.com/moisesPompilio/projeto_x/src/pkg/config/DB"
)

type Userusecase struct {
	UserRepositorie userrepositorie.IUserRepositorie
}

func NewUserusecase() IUserUsecase {
	repositorie := userrepositorie.NewUserRepository(db.GetReaderPostgreSQL(), db.GetWriterPostgreSQL())

	Userusecase := Userusecase{UserRepositorie: repositorie}
	return &Userusecase
}
