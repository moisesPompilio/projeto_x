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
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/output"
	usecase_mocks "github.com/moisesPompilio/projeto_x/src/pkg/mocks/usecase"
)

func setupTest() (token output.Token, login input.Login) {
	token = output.Token{Token: "dbkfsd", RefreshToken: "kjbdsdkjf"}

	login = input.Login{}

	return
}

func TestLogin_Valid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	token, login := setupTest()
	jsonBody, err := json.Marshal(login)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.NewReader(string(jsonBody))

	r := httptest.NewRequest(http.MethodPost, "/users/login", body)
	w := httptest.NewRecorder()

	MockIUserUsecase := usecase_mocks.NewMockIUserUsecase(mockCtrl)
	testUserUseCase := &userhandles.Userhandles{Usecase: MockIUserUsecase}

	MockIUserUsecase.EXPECT().LoginUseCase(login, gomock.Any()).Return(token, nil).Times(1)

	testUserUseCase.LoginHandle(w, r)

	got := w.Result()
	assert.Equal(t, got.StatusCode, http.StatusOK)

	var respBody output.Token
	err = json.NewDecoder(got.Body).Decode(&respBody)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, respBody, token)

}

func TestLogin_Invalid_MissingDataCorretFormat(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	token, _ := setupTest()
	login := "login"
	jsonBody, err := json.Marshal(login)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.NewReader(string(jsonBody))

	r := httptest.NewRequest(http.MethodPost, "/users/login", body)
	w := httptest.NewRecorder()

	MockIUserUsecase := usecase_mocks.NewMockIUserUsecase(mockCtrl)
	testUserUseCase := &userhandles.Userhandles{Usecase: MockIUserUsecase}

	MockIUserUsecase.EXPECT().LoginUseCase(gomock.Any(), gomock.Any()).Return(token, nil).Times(0)

	testUserUseCase.LoginHandle(w, r)

	got := w.Result()
	assert.Equal(t, got.StatusCode, http.StatusBadRequest)
}

func TestLogin_Invalid_ErrorInUserUseCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	token, login := setupTest()
	jsonBody, err := json.Marshal(login)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.NewReader(string(jsonBody))

	r := httptest.NewRequest(http.MethodPost, "/users/login", body)
	w := httptest.NewRecorder()

	MockIUserUsecase := usecase_mocks.NewMockIUserUsecase(mockCtrl)
	testUserUseCase := &userhandles.Userhandles{Usecase: MockIUserUsecase}

	MockIUserUsecase.EXPECT().LoginUseCase(login, gomock.Any()).Return(token, errors.New("falha ao fazer login")).Times(1)

	testUserUseCase.LoginHandle(w, r)

	got := w.Result()
	assert.Equal(t, got.StatusCode, http.StatusInternalServerError)
}
