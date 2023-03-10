package userrepositorie

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/moisesPompilio/projeto_x/src/internal/domain"
)

type ImplemetationUserRepositorieSQL struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewUserRepository(writer, reader *sqlx.DB) IUserRepositorie {
	return &ImplemetationUserRepositorieSQL{writer: writer, reader: reader}
}

func (repo *ImplemetationUserRepositorieSQL) Create(ctx context.Context, newUser domain.User) error {

	_, err := repo.writer.ExecContext(ctx, `INSERT INTO users (id, nick_name, name, email, password, created_at, updated_at) 
												VALUES ($1, $2, $3, $4, $5, $6, $7)`, newUser.ID, newUser.NickName, newUser.Name, newUser.Email, newUser.Password, newUser.CreatedAt, newUser.UpdatedAt)

	if err != nil {
		fmt.Println(err)
		return errors.New("usuário não cadastrado")
	}

	return nil
}

// func (repo *repoSqlx) GetByID(ctx context.Context, ID int64) (*domain.UserWithoutPassword, error) {

// 	var user domain.UserWithoutPassword

// 	err := repo.reader.GetContext(ctx, &user, `
// 		SELECT
// 			id,
// 			nick_name,
// 			name,
// 			email,
// 			created_at,
// 			updated_at
// 		FROM users
// 		WHERE id=?
// 	`, ID)

// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, errors.New("usuário não encontrado")
// 	}

// 	return &user, nil

// }

// func (repo *repoSqlx) GetAll(ctx context.Context) ([]domain.UserWithoutPassword, error) {

// 	var user []domain.UserWithoutPassword

// 	err := repo.reader.SelectContext(ctx, &user, `
// 		SELECT
// 			id,
// 			nick_name,
// 			name,
// 			email,
// 			created_at,
// 			updated_at
// 		FROM users
// 	`)

// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, errors.New("usuário não encontrado")
// 	}

// 	return user, nil

// }

func (repo *ImplemetationUserRepositorieSQL) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	err := repo.reader.GetContext(ctx, &user, `
		SELECT 
			id,
			nick_name,
			name,
			email,
			password,
			created_at,
			updated_at
		FROM users 
		WHERE email = $1
		LIMIT 1
	`, email)

	if err != nil {
		fmt.Println(err)
		return domain.User{}, errors.New("usuário não encontrado")
	}

	return user, nil
}

func (repo *ImplemetationUserRepositorieSQL) DeleteByID(ctx context.Context, ID string) error {
	_, err := repo.writer.ExecContext(ctx, `DELETE FROM users WHERE id = $1`, ID)

	if err != nil {
		fmt.Println(err)
		return errors.New("usuário não cadastrado")
	}

	return nil
}
