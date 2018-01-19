package api

import (
	"net/http"

	// Third party packages
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserSchema Struct
type UserSchema struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	UUID   string        `json:"uuid" bson:"uuid"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}

// Users resource struct definition
type Users struct {
	session *mgo.Session
}

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
		ID:     bson.NewObjectId(),
		Name:   "John Appleseed",
		Gender: "Age",
		Age:    35,
		UUID:   string(id),
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	render.JSON(res, req, user)
}

// Create creates a new user resource
func (rs Users) Create(res http.ResponseWriter, req *http.Request) {
	id, _ := uuid.NewV4()
	user := UserSchema{
		ID:     bson.NewObjectId(),
		UUID:   id.String(),
		Name:   "Susan Appleseed",
		Age:    33,
		Gender: "female",
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	render.JSON(res, req, user)
}
