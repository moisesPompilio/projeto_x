package userhandles

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

const (
	badRequestErrorMsg = `{"error": "Bad Request"}`
	contentTypeJSON    = "application/json"
)

func (usecase *Userhandles) GetByIDHandles(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, badRequestErrorMsg, http.StatusBadRequest)
		return
	}

	uuid, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, badRequestErrorMsg, http.StatusBadRequest)
		return
	}

	user, err := usecase.Usecase.GetUserByIdUseCase(context.Background(), uuid)
	if err != nil {
		w.Header().Set("Content-Type", contentTypeJSON)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", contentTypeJSON)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, badRequestErrorMsg, http.StatusInternalServerError)
		return
	}
}
