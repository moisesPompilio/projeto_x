package handles

import (
	"context"
	"encoding/json"
	"net/http"

	usecase "github.com/moisesPompilio/projeto_x/src/internal/UseCase"
	"github.com/moisesPompilio/projeto_x/src/internal/interfaces"
)

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	var login interfaces.Login
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := usecase.LoginUseCase(login, context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
