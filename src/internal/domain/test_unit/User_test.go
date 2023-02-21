package testunit

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moisesPompilio/projeto_x/src/internal/domain"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/input"
)

func TestNewUserValidInput(t *testing.T) {
	createUserDTO := input.CreateUserDTO{
		NickName: "test_nickname",
		Name:     "test_name",
		Email:    "test@email.com",
		Password: "test_password",
	}
	user, err := domain.NewUser(createUserDTO)
	assert.Nil(t, err)
	assert.Equal(t, createUserDTO.NickName, user.NickName)
	assert.Equal(t, createUserDTO.Name, user.Name)
	assert.Equal(t, createUserDTO.Email, user.Email)
	assert.Equal(t, createUserDTO.Password, user.Password)
	assert.NotNil(t, user.UpdatedAt)
	assert.NotNil(t, user.CreatedAt)
}

func TestNewUserMissingFields(t *testing.T) {
	createUserDTO := input.CreateUserDTO{
		NickName: "test_nickname",
		Password: "test_password",
	}
	_, err := domain.NewUser(createUserDTO)
	assert.NotNil(t, err)
}

func TestValid(t *testing.T) {
	user := domain.User{
		NickName: "test_nickname",
		Name:     "test_name",
		Email:    "test@email.com",
		Password: "test_password",
	}
	err := user.Valid()
	assert.Nil(t, err)
}

func TestIsValidMissingNickName(t *testing.T) {
	user := domain.User{
		Name:     "test_name",
		Email:    "test@email.com",
		Password: "test_password",
	}
	err := user.Valid()

	assert.Equal(t, "field \"NickName\" cannot be empty", err.Error())
}

func TestIsValidMissingName(t *testing.T) {
	user := domain.User{
		NickName: "test_nickname",
		Email:    "test@email.com",
		Password: "test_password",
	}
	err := user.Valid()
	assert.Equal(t, "field \"Name\" cannot be empty", err.Error())
}

func TestIsValidMissingEmail(t *testing.T) {
	user := domain.User{
		NickName: "test_nickname",
		Name:     "test_name",
		Password: "test_password",
	}
	err := user.Valid()
	assert.Equal(t, "field \"Email\" cannot be empty", err.Error())
}

func TestIsValidMissingPassword(t *testing.T) {
	user := domain.User{
		NickName: "test_nickname",
		Name:     "test_name",
		Email:    "test@email.com",
	}
	err := user.Valid()
	assert.Equal(t, "field \"Password\" cannot be empty", err.Error())
}
