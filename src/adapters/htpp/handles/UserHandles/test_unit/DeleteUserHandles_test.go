package testunit

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	userhandles "github.com/moisesPompilio/projeto_x/src/adapters/htpp/handles/UserHandles"
	usecase_mocks "github.com/moisesPompilio/projeto_x/src/pkg/mocks/usecase"
)

func TestDeleteUserHandles_Valid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	id := "asjk"

	r := httptest.NewRequest(http.MethodDelete, ("/users?id=" + id), nil)
	w := httptest.NewRecorder()

	MockIUserUsecase := usecase_mocks.NewMockIUserUsecase(mockCtrl)
	testUserUseCase := &userhandles.Userhandles{Usecase: MockIUserUsecase}

	MockIUserUsecase.EXPECT().DeleteUsersById(id, gomock.Any()).Return(nil).Times(1)

	testUserUseCase.DeleteUserHandles(w, r)

	got := w.Result()
	assert.Equal(t, http.StatusOK, got.StatusCode)
}

func TestDeleteUserHandles_Invalid_MissingDataCorrectFormat(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	r := httptest.NewRequest(http.MethodDelete, "/users/", nil)
	w := httptest.NewRecorder()

	MockIUserUsecase := usecase_mocks.NewMockIUserUsecase(mockCtrl)
	testUserUseCase := &userhandles.Userhandles{Usecase: MockIUserUsecase}

	MockIUserUsecase.EXPECT().DeleteUsersById(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	testUserUseCase.DeleteUserHandles(w, r)

	got := w.Result()
	assert.Equal(t, http.StatusBadRequest, got.StatusCode)
}

func TestDeleteUserHandles_Invalid_ErrorInUserUseCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	id := "idUser"

	r := httptest.NewRequest(http.MethodDelete, ("/users?id=" + id), nil)
	w := httptest.NewRecorder()

	MockIUserUsecase := usecase_mocks.NewMockIUserUsecase(mockCtrl)
	testUserUseCase := &userhandles.Userhandles{Usecase: MockIUserUsecase}

	MockIUserUsecase.EXPECT().DeleteUsersById(id, gomock.Any()).Return(errors.New("failed to delete user")).Times(1)

	testUserUseCase.DeleteUserHandles(w, r)

	got := w.Result()
	assert.Equal(t, http.StatusInternalServerError, got.StatusCode)
}
