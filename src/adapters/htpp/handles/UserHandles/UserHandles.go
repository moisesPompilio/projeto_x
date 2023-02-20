package userhandles

import (
	userusecase "github.com/moisesPompilio/projeto_x/src/internal/UseCase/UserUseCase"
)

type Userhandles struct {
	Usecase userusecase.Userusecase
}

func NewUserhandles() *Userhandles {
	usecase := userusecase.NewUserusecase()
	return &Userhandles{Usecase: usecase}
}
