package userusecase

import (
	"context"
)

func (usecase *Userusecase) DeleteUsersById(data string, ctx context.Context) error {
	err := usecase.UserRepositorie.DeleteByID(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
