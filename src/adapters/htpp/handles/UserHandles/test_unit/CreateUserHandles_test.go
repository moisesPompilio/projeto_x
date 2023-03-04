package testunit

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	userhandles "github.com/moisesPompilio/projeto_x/src/adapters/htpp/handles/UserHandles"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/input"
	usecase_mocks "github.com/moisesPompilio/projeto_x/src/pkg/mocks/usecase"
)

func setupTestCreateUser() (createUserInvalid string, createUserValid input.CreateUserDTO) {
	createUserInvalid = "teste"

	createUserValid = input.CreateUserDTO{}
	return
}

func TestCreateUserHandle_Valid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	_, createUserDTO := setupTestCreateUser()
	jsonBody, err := json.Marshal(createUserDTO)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.NewReader(string(jsonBody))

	r := httptest.NewRequest(http.MethodPost, "/users", body)
	w := httptest.NewRecorder()

	MockIUserUsecase := usecase_mocks.NewMockIUserUsecase(mockCtrl)
	testUserUseCase := &userhandles.Userhandles{Usecase: MockIUserUsecase}

	MockIUserUsecase.EXPECT().CreateUserUseCase(createUserDTO, gomock.Any()).Return(nil).Times(1)

	testUserUseCase.CreateUserHandle(w, r)

	got := w.Result()
	assert.Equal(t, got.StatusCode, http.StatusCreated)

}

func TestCreateUserHandle_Invalid_MissingDataCOrretFormat(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	createUserDTO, _ := setupTestCreateUser()
	jsonBody, err := json.Marshal(createUserDTO)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.NewReader(string(jsonBody))

	r := httptest.NewRequest(http.MethodPost, "/users/login", body)
	w := httptest.NewRecorder()

	MockIUserUsecase := usecase_mocks.NewMockIUserUsecase(mockCtrl)
	testUserUseCase := &userhandles.Userhandles{Usecase: MockIUserUsecase}

	MockIUserUsecase.EXPECT().CreateUserUseCase(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	testUserUseCase.CreateUserHandle(w, r)

	got := w.Result()
	assert.Equal(t, got.StatusCode, http.StatusBadRequest)
}

func TestCreateUserHandle_Invalid_ErrorInUserUseCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	_, login := setupTestCreateUser()
	jsonBody, err := json.Marshal(login)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.NewReader(string(jsonBody))

	r := httptest.NewRequest(http.MethodPost, "/users/login", body)
	w := httptest.NewRecorder()

	MockIUserUsecase := usecase_mocks.NewMockIUserUsecase(mockCtrl)
	testUserUseCase := &userhandles.Userhandles{Usecase: MockIUserUsecase}

	MockIUserUsecase.EXPECT().CreateUserUseCase(login, gomock.Any()).Return(errors.New("falha ao fazer login")).Times(1)

	testUserUseCase.CreateUserHandle(w, r)

	got := w.Result()
	assert.Equal(t, got.StatusCode, http.StatusInternalServerError)
}
