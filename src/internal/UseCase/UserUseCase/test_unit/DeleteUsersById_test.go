package testunit

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"

// 	userusecase "github.com/moisesPompilio/projeto_x/src/internal/UseCase/UserUseCase"
// 	mocks "github.com/moisesPompilio/projeto_x/src/pkg/mocks/repositorie"
// )

// func TestDeleteUsersById_Valid(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockUserRepositorie := mocks.NewMockIUserRepositorie(mockCtrl)
// 	testUserUseCase := &userusecase.Userusecase{UserRepositorie: mockUserRepositorie}

// 	context := context.Background()
// 	id := "idTest"

// 	mockUserRepositorie.EXPECT().DeleteByID(context, id).Return(nil).Times(1)

// 	err := testUserUseCase.DeleteUsersById(id, context)

// 	assert.Nil(t, err)

// }

// func TestDeleteUsersById_Invalid_IDUser(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockUserRepositorie := mocks.NewMockIUserRepositorie(mockCtrl)
// 	testUserUseCase := &userusecase.Userusecase{UserRepositorie: mockUserRepositorie}

// 	context := context.Background()
// 	id := ""

// 	mockUserRepositorie.EXPECT().DeleteByID(context, gomock.Any()).Return(nil).Times(0)

// 	err := testUserUseCase.DeleteUsersById(id, context)

// 	assert.NotNil(t, err)
// 	assert.Equal(t, "id invalid", err.Error())
// }

// func TestDeleteUsersById_Invalid_ErrorRepositorie(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockUserRepositorie := mocks.NewMockIUserRepositorie(mockCtrl)
// 	testUserUseCase := &userusecase.Userusecase{UserRepositorie: mockUserRepositorie}

// 	context := context.Background()
// 	id := "idTest"

// 	mockUserRepositorie.EXPECT().DeleteByID(context, id).Return(errors.New("falha ao fazer o teste")).Times(1)

// 	err := testUserUseCase.DeleteUsersById(id, context)

// 	assert.NotNil(t, err)
// 	assert.Equal(t, "falha ao fazer o teste", err.Error())
// }
