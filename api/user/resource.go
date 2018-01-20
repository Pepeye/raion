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

// Create creates a new user resource
func (rs Resource) Create(w http.ResponseWriter, r *http.Request) {
	id := uuid.NewV4()
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

// List retrieves a list of user resources
func (rs Resource) List(w http.ResponseWriter, r *http.Request) {
	// set response header once
	w.Header().Set("Content-Type", "application/json")
	users := []Schema{}

	if err := rs.Session.DB("raion").C("users").Find(bson.M{}).All(&users); err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, map[string]string{"message": "unable to find users"})
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, users)
}

// func validateObjectID(idParamStr string) (oid bson.ObjectId, err map[string]string) {
// 	if !bson.IsObjectIdHex(idParamStr) {
// 		err = map[string]string{"message": "invalid :id parameter provided"}
// 		return
// 	}
// 	oid = bson.ObjectIdHex(idParamStr)
// 	return
// }

func validateObjectID(idParamStr string, w http.ResponseWriter, r *http.Request) (oid bson.ObjectId) {
	if !bson.IsObjectIdHex(idParamStr) {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"message": "invalid :id parameter provided"})
		return
	}
	oid = bson.ObjectIdHex(idParamStr)
	return
}

// Get retrieves an individual user resource
func (rs Resource) Get(w http.ResponseWriter, r *http.Request) {
	// set response header once
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")
	oid := validateObjectID(id, w, r)
	user := Schema{}

	if err := rs.Session.DB("raion").C("users").FindId(oid).One(&user); err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, map[string]string{"message": "user not found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, user)
}

// Update updates an individual user resource
func (rs Resource) Update(w http.ResponseWriter, r *http.Request) {
	// set response header once
	w.Header().Set("Content-Type", "application/json")

	user := Schema{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"message": "no updated user data was sent"})
		return
	}

	id := chi.URLParam(r, "id")
	oid := validateObjectID(id, w, r)
	query := bson.M{"_id": oid}
	data := bson.M{
		"$set": bson.M{
			"name":   user.Name,
			"gender": user.Gender,
			"age":    user.Age,
		},
	}

	if err := rs.Session.DB("raion").C("users").Update(query, data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"message": "unable to update user"})
		return
	}
	// get modified user from database
	rs.Session.DB("raion").C("users").FindId(oid).One(&user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, user)
}

// Delete removes an individual user resource
func (rs Resource) Delete(w http.ResponseWriter, r *http.Request) {
	// set response type once
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")
	oid := validateObjectID(id, w, r)
	if err := rs.Session.DB("raion").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"message": "unable to delete user from database"})
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]string{"message": "user deleted from database", "deleteduserid": id})
}
