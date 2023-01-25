package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/moisesPompilio/projeto_x/src/internal/interfaces"
)

type User struct {
	ID        uuid.UUID  `db:"id"`
	NickName  string     `db:"nick_name"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

func NewUser(createUserDTO interfaces.CreateUserDTO) (User, error) {
	now := time.Now()
	ID := uuid.New()
	var user User = User{
		ID:        ID,
		NickName:  createUserDTO.NickName,
		Name:      createUserDTO.Name,
		Email:     createUserDTO.Email,
		Password:  createUserDTO.Password,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	erro := user.valid()
	if erro != nil {
		return user, erro
	}
	return user, nil
}

func (user *User) valid() error {
	if user.Email == "" || user.Name == "" || user.Password == "" || user.NickName == "" {
		return errors.New("user invalid")
	}
	return nil
}
