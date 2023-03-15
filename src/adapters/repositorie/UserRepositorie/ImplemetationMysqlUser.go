package userrepositorie

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/moisesPompilio/projeto_x/src/internal/domain"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/output"
)

type ImplemetationUserRepositorieSQL struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewUserRepository(writer, reader *sqlx.DB) IUserRepositorie {
	return &ImplemetationUserRepositorieSQL{writer: writer, reader: reader}
}

func (repo *ImplemetationUserRepositorieSQL) Create(ctx context.Context, newUser domain.User) error {
	stmt, err := repo.writer.PrepareContext(ctx, `INSERT INTO users (id, nick_name, name, email, roles, password, created_at, updated_at) 
                                                VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)
	if err != nil {
		return errors.New("failed to prepare query: " + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, newUser.ID, newUser.NickName, newUser.Name, newUser.Email, newUser.Roles, newUser.Password, newUser.CreatedAt, newUser.UpdatedAt)
	if err != nil {
		return errors.New("failed to execute query: " + err.Error())
	}

	return nil
}

func (repo *ImplemetationUserRepositorieSQL) GetByID(ctx context.Context, ID uuid.UUID) (*output.UserOutpDTO, error) {
	var userOutpDTO output.UserOutpDTO

	err := repo.reader.GetContext(ctx, &userOutpDTO, `
        SELECT
            id,
            nick_name,
            name,
            email
        FROM
            users
        WHERE
            id = $1
    `, ID)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("usuário não encontrado. Motivo: " + err.Error())
	}

	return &userOutpDTO, nil
}

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
