package testunit

import (
	"testing"

	"github.com/moisesPompilio/projeto_x/src/internal/domain"
)

func TestValid(t *testing.T) {
	user := domain.User{
		NickName: "test_nickname",
		Name:     "test_name",
		Email:    "test@email.com",
		Password: "test_password",
	}
	err := user.Valid()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestIsValidMissingNickName(t *testing.T) {
	user := domain.User{
		Name:     "test_name",
		Email:    "test@email.com",
		Password: "test_password",
	}
	err := user.Valid()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestIsValidMissingName(t *testing.T) {
	user := domain.User{
		NickName: "test_nickname",
		Email:    "test@email.com",
		Password: "test_password",
	}
	err := user.Valid()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestIsValidMissingEmail(t *testing.T) {
	user := domain.User{
		NickName: "test_nickname",
		Name:     "test_name",
		Password: "test_password",
	}
	err := user.Valid()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestIsValidMissingPassword(t *testing.T) {
	user := domain.User{
		NickName: "test_nickname",
		Name:     "test_name",
		Email:    "test@email.com",
	}
	err := user.Valid()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
