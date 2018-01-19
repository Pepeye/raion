package user

import (
	// Third party packages
	"github.com/go-chi/chi"
)

// Routes creates a REST router for user resource
func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()

	// add middleware specific to user Routes

	r.Post("/", rs.Create) // POST /users
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get) // GET /users/{id}
	})

	return r
}
