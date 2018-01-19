package main

import (
	// standard library packages
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	// third party packages
	"github.com/Pepeye/raion/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

const appname string = "raion"
const host string = "localhost"
const port string = ":3333"

// Message struct
type Message struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Server    `json:"server,omitempty"`
	CreatedAt time.Time `json:"createdat"`
}

// Server struct
type Server struct {
	App  string `json:"app"`
	Host string `json:"host"`
	Port string `json:"port"`
}

func main() {
	// new router
	app := chi.NewRouter()

	// middleware
	app.Use(middleware.RequestID)
	app.Use(middleware.RealIP)
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)
	app.Use(middleware.URLFormat)
	app.Use(render.SetContentType(render.ContentTypeJSON))

	// public routes
	app.Get("/", handlerFn)
	app.Get("/user/{id}", getUser)

	// start server
	// defer http.ListenAndServe(":3001", app)
	fmt.Printf("[%s]: Server running...", appname)
	err := http.ListenAndServe(host+port, app)
	if err != nil {
		log.Fatal("Serving: ", err)
	}
}

func handlerFn(res http.ResponseWriter, req *http.Request) {
	// create a new message
	msg := Message{
		ID:        "f80b342c-f90c-4804-9df1-faeb244ab9b8",
		Message:   "raion api",
		Server:    Server{appname, host, port},
		CreatedAt: time.Now(),
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(msg)
}

func getUser(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	user := models.User{
		Name:   "Goose Gander",
		Gender: "Age",
		Age:    35,
		ID:     string(id),
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	render.JSON(res, req, user)
}
