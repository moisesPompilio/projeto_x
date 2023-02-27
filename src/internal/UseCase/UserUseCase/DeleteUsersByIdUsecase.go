package userusecase

import (
	"context"
	"errors"
)

func (usecase *Userusecase) DeleteUsersById(data string, ctx context.Context) error {
	if data == "" {
		return errors.New("id invalid")
	}

	err := usecase.UserRepositorie.DeleteByID(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
