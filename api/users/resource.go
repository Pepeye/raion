package users

import (
	"net/http"

	// Third party packages
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Resource struct for users
type Resource struct {
	session *mgo.Session
}

// Get retrieves an individual user resource
func (rs Resource) Get(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	user := Schema{
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
func (rs Resource) Create(res http.ResponseWriter, req *http.Request) {
	id, _ := uuid.NewV4()
	user := Schema{
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
