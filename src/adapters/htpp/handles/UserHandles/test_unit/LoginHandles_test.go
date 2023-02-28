package testunit

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"

	userhandles "github.com/moisesPompilio/projeto_x/src/adapters/htpp/handles/UserHandles"
	"github.com/moisesPompilio/projeto_x/src/internal/domain"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/input"
)

func variblesTest() (user domain.User, userErr domain.User, login input.Login) {
	now := time.Now()
	user = domain.User{
		ID:        uuid.MustParse("0d3fd34a-98cc-11eb-8ac8-0242ac130003"),
		NickName:  "janedoe",
		Name:      "Jane Doe",
		Email:     "johndoe@example.com",
		Password:  "testpassword",
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	userErr = domain.User{
		ID:        uuid.MustParse("f0b72280-9935-11eb-8ac8-0242ac130003"),
		NickName:  "johndoe",
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		Password:  "wrongPassword",
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	login = input.Login{
		Email:    "johndoe@example.com",
		Password: "testpassword",
	}

	return
}
func TestLogin_Valid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MockIUserUsecase := usecase_mocks.NewMockIUserRepositorie(mockCtrl)

	testUserUseCase := &userhandles.Userhandles{Usecase: MockIUserUsecase}

	// context := context.Background()
	// user, _, login := variblesTest()

	// mockUserRepositorie.EXPECT().GetUserByEmail(context, login.Email).Return(user, nil).Times(1)

	// token, err := testUserUseCase.LoginUseCase(login, context)

	// assert.Nil(t, err)
	// assert.NotNil(t, token)

}

func TestLogin_MissingUser(t *testing.T) {
	// mockCtrl := gomock.NewController(t)
	// defer mockCtrl.Finish()

	// mockUserRepositorie := mocks
	// testUserUseCase := &userusecase.Userusecase{UserRepositorie: mockUserRepositorie}

	// context := context.Background()
	// _, _, login := variblesTest()

	// mockUserRepositorie.EXPECT().GetUserByEmail(context, login.Email).Return(domain.User{}, errors.New("usuário não encontrado")).Times(1)

	// _, err := testUserUseCase.LoginUseCase(login, context)

	// assert.NotNil(t, err)
	// assert.Equal(t, "usuário não encontrado", err.Error())
}
func TestLogin_InvalidUser(t *testing.T) {
	// mockCtrl := gomock.NewController(t)
	// defer mockCtrl.Finish()

	// mockUserRepositorie := mocks.NewMockIUserRepositorie(mockCtrl)
	// testUserUseCase := &userusecase.Userusecase{UserRepositorie: mockUserRepositorie}

	// context := context.Background()
	// _, userErr, login := variblesTest()

	// mockUserRepositorie.EXPECT().GetUserByEmail(context, login.Email).Return(userErr, nil).Times(1)

	// _, err := testUserUseCase.LoginUseCase(login, context)

	// assert.NotNil(t, err)
	// assert.Equal(t, "Usuario invalido!", err.Error())

}