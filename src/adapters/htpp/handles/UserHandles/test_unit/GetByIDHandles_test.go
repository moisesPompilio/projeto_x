package testunit

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	userhandles "github.com/moisesPompilio/projeto_x/src/adapters/htpp/handles/UserHandles"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/output"
	usecase_mocks "github.com/moisesPompilio/projeto_x/src/pkg/mocks/usecase"
)

func setupGetByIDHandles() (user_outp output.UserOutpDTO) {
	user_outp = output.UserOutpDTO{
		Id:       uuid.MustParse("0d3fd34a-98cc-11eb-8ac8-0242ac130003"),
		NickName: "janedoe",
		Name:     "Jane Doe",
		Email:    "johndoe@example.com",
	}

	return
}

func TestGetByIDHandles_Valid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user_valid := setupGetByIDHandles()
	id := user_valid.Id.String()

	r := httptest.NewRequest(http.MethodGet, ("/users?id=" + id), nil)
	w := httptest.NewRecorder()

	MockIUserUsecase := usecase_mocks.NewMockIUserUsecase(mockCtrl)
	testUserUseCase := &userhandles.Userhandles{Usecase: MockIUserUsecase}

	MockIUserUsecase.EXPECT().GetUserByIdUseCase(gomock.Any(), user_valid.Id).Return(&user_valid, nil).Times(1)

	testUserUseCase.GetByIDHandles(w, r)

	got := w.Result()
	assert.Equal(t, http.StatusOK, got.StatusCode)
}

func TestGetByIDHandles_Invalid_MissingDataCorrectFormat(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	r := httptest.NewRequest(http.MethodGet, "/users/", nil)
	w := httptest.NewRecorder()

	MockIUserUsecase := usecase_mocks.NewMockIUserUsecase(mockCtrl)
	testUserUseCase := &userhandles.Userhandles{Usecase: MockIUserUsecase}

	MockIUserUsecase.EXPECT().GetUserByIdUseCase(gomock.Any(), gomock.Any()).Return(nil, nil).Times(0)

	testUserUseCase.GetByIDHandles(w, r)

	got := w.Result()
	assert.Equal(t, http.StatusBadRequest, got.StatusCode)
}

func TestGetByIDHandles_Invalid_ErrorInUserUseCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user_valid := setupGetByIDHandles()
	id := user_valid.Id.String()

	r := httptest.NewRequest(http.MethodGet, ("/users?id=" + id), nil)
	w := httptest.NewRecorder()

	MockIUserUsecase := usecase_mocks.NewMockIUserUsecase(mockCtrl)
	testUserUseCase := &userhandles.Userhandles{Usecase: MockIUserUsecase}

	MockIUserUsecase.EXPECT().GetUserByIdUseCase(gomock.Any(), user_valid.Id).Return(nil, errors.New("failed to get user")).Times(1).Times(1)

	testUserUseCase.GetByIDHandles(w, r)

	got := w.Result()
	assert.Equal(t, http.StatusInternalServerError, got.StatusCode)
}
