package routes

import (
	"github.com/go-chi/chi/v5"

	userroutes "github.com/moisesPompilio/projeto_x/src/adapters/htpp/routes/UserRoutes"
)

func LoadRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/users", userroutes.UserRoutes())
	return router
}
