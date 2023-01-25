package usecase

import (
	"context"

	userrepositorie "github.com/moisesPompilio/projeto_x/src/adapters/repositorie/UserRepositorie"
	db "github.com/moisesPompilio/projeto_x/src/pkg/config/DB"
)

func DeleteUsersById(data string, ctx context.Context) error {
	conection := userrepositorie.NewUserRepository(db.GetReaderPostgreSQL(), db.GetWriterPostgreSQL())
	err := conection.DeleteByID(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
