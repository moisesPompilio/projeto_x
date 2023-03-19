package userhandles

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/input"
)

const (
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json"
	invalidMethodErr  = "Invalid method"
)

func (usecase *Userhandles) LoginHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, invalidMethodErr, http.StatusMethodNotAllowed)
		return
	}

	var login input.Login
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := usecase.Usecase.LoginUseCase(login, context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(contentTypeHeader, jsonContentType)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
