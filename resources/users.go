package resources

import (
	"net/http"

	// Third party packages
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	uuid "github.com/satori/go.uuid"
)

// UserSchema Struct
type UserSchema struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

// Users resource struct definition
type Users struct{}

// Routes creates a REST router for user resource
func (rs Users) Routes() chi.Router {
	r := chi.NewRouter()

	// add middleware specific to user Routes

	r.Post("/", rs.Create) // POST /users
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get) // GET /users/{id}
	})

	return r
}

// Get retrieves an individual user resource
func (rs Users) Get(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	user := UserSchema{
		Name:   "John Appleseed",
		Gender: "Age",
		Age:    35,
		ID:     string(id),
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	render.JSON(res, req, user)
}

// Create creates a new user resource
func (rs Users) Create(res http.ResponseWriter, req *http.Request) {
	id, _ := uuid.NewV4()
	user := UserSchema{
		ID:     id.String(),
		Name:   "Susan Appleseed",
		Age:    33,
		Gender: "female",
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	render.JSON(res, req, user)
}
