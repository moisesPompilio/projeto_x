package testunit

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"

// 	userusecase "github.com/moisesPompilio/projeto_x/src/internal/UseCase/UserUseCase"
// 	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/input"
// 	mocks "github.com/moisesPompilio/projeto_x/src/pkg/mocks/repositorie"
// )

// func variblesTest() (createDTOInvalid input.CreateUserDTO, createUserDTO input.CreateUserDTO) {

// 	createDTOInvalid = input.CreateUserDTO{
// 		NickName: "",
// 		Name:     "",
// 		Email:    "",
// 		Password: "",
// 	}

// 	createUserDTO = input.CreateUserDTO{
// 		NickName: "janedoe",
// 		Name:     "Jane Doe",
// 		Email:    "johndoe@example.com",
// 		Password: "testpassword",
// 	}

// 	return
// }
// func TestCreateUserUseCase_Valid(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockUserRepositorie := mocks.NewMockIUserRepositorie(mockCtrl)
// 	testUserUseCase := &userusecase.Userusecase{UserRepositorie: mockUserRepositorie}

// 	context := context.Background()
// 	_, createUserDTO := variblesTest()

// 	mockUserRepositorie.EXPECT().Create(context, gomock.Any()).Return(nil).Times(1)

// 	err := testUserUseCase.CreateUserUseCase(createUserDTO, context)

// 	assert.Nil(t, err)

// }

// func TestLogin_InvalidUser(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockUserRepositorie := mocks.NewMockIUserRepositorie(mockCtrl)
// 	testUserUseCase := &userusecase.Userusecase{UserRepositorie: mockUserRepositorie}

// 	context := context.Background()
// 	createUserDTO, _ := variblesTest()

// 	mockUserRepositorie.EXPECT().Create(context, gomock.Any()).Return(nil).Times(0)

// 	err := testUserUseCase.CreateUserUseCase(createUserDTO, context)

// 	assert.NotNil(t, err)
// }

// func TestLogin_FailedInsertUser(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockUserRepositorie := mocks.NewMockIUserRepositorie(mockCtrl)
// 	testUserUseCase := &userusecase.Userusecase{UserRepositorie: mockUserRepositorie}

// 	context := context.Background()
// 	_, createUserDTO := variblesTest()

// 	mockUserRepositorie.EXPECT().Create(context, gomock.Any()).Return(errors.New("falha ao fazer o teste")).Times(1)

// 	err := testUserUseCase.CreateUserUseCase(createUserDTO, context)

// 	assert.NotNil(t, err)
// 	assert.Equal(t, "falha ao fazer o teste", err.Error())
// }
