package userhandles

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/input"
)

func (usecase *Userhandles) LoginHandle(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
