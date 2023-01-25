package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/moisesPompilio/projeto_x/src/adapters/htpp/handles"
)

func UserRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", handles.CreateUserHandle)
	router.Post("/login", handles.LoginHandle)
	router.Delete("/{id}", handles.DeleteUserHandles)
	return router
}
