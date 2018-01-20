package user

import (
	"encoding/json"
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
	Session *mgo.Session
}

// Get retrieves an individual user resource
func (rs Resource) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// set response header once
	w.Header().Set("Content-Type", "application/json")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"message": "invalid :id parameter provided"})
		return
	}

	oid := bson.ObjectIdHex(id)
	user := Schema{}

	if err := rs.Session.DB("raion").C("users").FindId(oid).One(&user); err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, map[string]string{"message": "user not found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, user)
}

// Create creates a new user resource
func (rs Resource) Create(w http.ResponseWriter, r *http.Request) {
	id, _ := uuid.NewV4()
	user := Schema{
		ID:   bson.NewObjectId(),
		UUID: id.String(),
	}

	json.NewDecoder(r.Body).Decode(&user)
	// user.UUID = id.String()
	db := rs.Session.Copy()
	defer db.Close()

	// insert user data
	err := db.DB("raion").C("users").Insert(user)

	if err != nil {
		if mgo.IsDup(err) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"message": "user with this Id already exists"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]error{"message": err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, user)
}
