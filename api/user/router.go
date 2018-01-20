package user

import (
	// Third party packages
	"github.com/go-chi/chi"
)

// Routes creates a REST router for user resource
func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()

	// add middleware specific to user Routes
	r.Get("/", rs.List)    // GET /users - read list of users
	r.Post("/", rs.Create) // POST /users - create new users
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get)    // GET /users/{id} - read a single user
		r.Put("/", rs.Update) // PUT /users/{id} - update asingle user
		// r.Delete("/", rs.Delete) // DELETE /users/{id} - delete a single user
	})

	return r
}
