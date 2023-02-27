package userhandles

import (
	userusecase "github.com/moisesPompilio/projeto_x/src/internal/UseCase/UserUseCase"
)

type Userhandles struct {
	Usecase userusecase.IUserUsecase
}

func NewUserhandles() IUserHandles {
	usecase := userusecase.NewUserusecase()
	return &Userhandles{Usecase: usecase}
}
