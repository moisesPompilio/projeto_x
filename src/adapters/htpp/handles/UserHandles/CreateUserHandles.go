package userhandles

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/input"
)

func (usecase *Userhandles) CreateUserHandle(w http.ResponseWriter, r *http.Request) {
	var createUserDTO input.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&createUserDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := usecase.Usecase.CreateUserUseCase(createUserDTO, context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w)
}
