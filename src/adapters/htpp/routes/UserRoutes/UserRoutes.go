package userroutes

import (
	"github.com/go-chi/chi/v5"

	handles "github.com/moisesPompilio/projeto_x/src/adapters/htpp/handles/CreateUser"
)

func UserRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", handles.CreateUserHandle)
	return router
}
