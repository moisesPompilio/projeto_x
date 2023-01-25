package handles

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	usecase "github.com/moisesPompilio/projeto_x/src/internal/UseCase"
)

func DeleteUserHandles(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	if userID == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}
	err := usecase.DeleteUsersById(userID, context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w)
}
