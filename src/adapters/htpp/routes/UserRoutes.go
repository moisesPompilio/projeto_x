package routes

import (
	"github.com/go-chi/chi/v5"

	userhandles "github.com/moisesPompilio/projeto_x/src/adapters/htpp/handles/UserHandles"
)

func UserRoutes() *chi.Mux {
	handler := userhandles.NewUserhandles()
	router := chi.NewRouter()
	router.Post("/", handler.CreateUserHandle)
	router.Post("/login", handler.LoginHandle)
	router.Delete("/{id}", handler.DeleteUserHandles)
	return router
}
