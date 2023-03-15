package userusecase

import (
	"context"
	"errors"
)

func (usecase *Userusecase) DeleteUsersById(id string, ctx context.Context) error {
	if id == "" {
		return errors.New("id invalid")
	}

	err := usecase.UserRepositorie.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
