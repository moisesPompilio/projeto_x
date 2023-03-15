package domain

import (
	"time"

	"github.com/google/uuid"

	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/input"
	validation "github.com/moisesPompilio/projeto_x/src/pkg/helps/Validation"
)

type User struct {
	ID        uuid.UUID  `db:"id"`
	NickName  string     `db:"nick_name"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Roles     string     `db:"roles"`
	Password  string     `db:"password"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

func NewUser(createUserDTO input.CreateUserDTO) (User, error) {
	now := time.Now()
	ID := uuid.New()
	var user User = User{
		ID:        ID,
		NickName:  createUserDTO.NickName,
		Name:      createUserDTO.Name,
		Email:     createUserDTO.Email,
		Roles:     createUserDTO.Roles,
		Password:  createUserDTO.Password,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	erro := user.Valid()
	if erro != nil {
		return user, erro
	}
	return user, nil
}

func (user *User) Valid() error {
	requiredFields := []string{"Email", "Name", "NickName", "Password"}
	err := validation.ValidateRequiredFields_Interfaces_Strings_Pointer(*user, requiredFields)
	return err
}
