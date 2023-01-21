package handles

import (
	"context"
	"encoding/json"
	"net/http"

	usecase "github.com/moisesPompilio/projeto_x/src/internal/UseCase/CreateUser"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces"
)

func CreateUserHandle(w http.ResponseWriter, r *http.Request) {
	var createUserDTO interfaces.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&createUserDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := usecase.UserUseCase(createUserDTO, context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
