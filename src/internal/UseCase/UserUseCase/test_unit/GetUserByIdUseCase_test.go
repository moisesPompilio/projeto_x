package testunit

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	userusecase "github.com/moisesPompilio/projeto_x/src/internal/UseCase/UserUseCase"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/output"
	mocks "github.com/moisesPompilio/projeto_x/src/pkg/mocks/repositorie"
)

func setupGetUserByIdUseCase() (user_outp output.UserOutpDTO) {
	user_outp = output.UserOutpDTO{
		Id:       uuid.MustParse("0d3fd34a-98cc-11eb-8ac8-0242ac130003"),
		NickName: "janedoe",
		Name:     "Jane Doe",
		Email:    "johndoe@example.com",
	}

	return
}

func TestGetUserByIdUseCase_Valid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepositorie := mocks.NewMockIUserRepositorie(mockCtrl)
	testUserUseCase := &userusecase.Userusecase{UserRepositorie: mockUserRepositorie}

	context := context.Background()
	user_outp := setupGetUserByIdUseCase()
	id := user_outp.Id

	mockUserRepositorie.EXPECT().GetByID(context, id).Return(&user_outp, nil).Times(1)

	user, err := testUserUseCase.GetUserByIdUseCase(context, id)

	assert.NotNil(t, user)
	assert.Equal(t, user, &user_outp)
	assert.Nil(t, err)
}

func TestGetUserByIdUseCase_Invalid_ErrorRepositorie(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepositorie := mocks.NewMockIUserRepositorie(mockCtrl)
	testUserUseCase := &userusecase.Userusecase{UserRepositorie: mockUserRepositorie}

	context := context.Background()
	id := uuid.UUID{}

	mockUserRepositorie.EXPECT().GetByID(context, id).Return(nil, errors.New("falha ao fazer o teste")).Times(1)

	_, err := testUserUseCase.GetUserByIdUseCase(context, id)

	assert.NotNil(t, err)
	assert.Equal(t, "falha ao fazer o teste", err.Error())
}
