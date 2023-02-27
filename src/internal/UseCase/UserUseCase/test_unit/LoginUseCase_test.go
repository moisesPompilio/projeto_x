package testunit_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	userusecase "github.com/moisesPompilio/projeto_x/src/internal/UseCase/UserUseCase"
	"github.com/moisesPompilio/projeto_x/src/internal/domain"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/input"
	mocks "github.com/moisesPompilio/projeto_x/src/pkg/mocks/adapters/repositorie"
)

func Test_User_NewUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepositorie := mocks.NewMockIUserRepositorie(mockCtrl)
	testUserUseCase := &userusecase.Userusecase{UserRepositorie: mockUserRepositorie}
	context := context.Background()
	mockUserRepositorie.EXPECT().GetUserByEmail(context, "testing@example.com").Return(domain.User{}, errors.New("usuário não encontrado")).Times(1)

	login := input.Login{
		Email:    "testing@example.com",
		Password: "testpassword",
	}

	_, err := testUserUseCase.LoginUseCase(login, context)
	if err == nil {
		t.Fatalf("expected no error, but got %v", err)
	}

}

// package User_test

// import (
// 	"testing"

// 	"github.com/golang/mock/gomock"

// 	"github.com/moisesPompilio/teste_mock_estudo/User"
// 	"github.com/moisesPompilio/teste_mock_estudo/mocks"
// )

// func Test_User_NewUser(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockUser := mocks.NewMockIUserInterface(mockCtrl)
// 	testUser := &User.User{IUser: mockUser}

// 	mockUser.EXPECT().AddUser(1, "foo").Return(nil).Times(1)
// 	testUser.User()
// }
