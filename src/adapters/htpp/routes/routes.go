package routes

import (
	"github.com/go-chi/chi/v5"
)

func LoadRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/users", UserRoutes())
	return router
}
