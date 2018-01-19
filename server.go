package main

import (
	// standard library packages
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	// third party packages

	"github.com/Pepeye/raion/resources"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	mgo "gopkg.in/mgo.v2"
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
	// app.Get("/user/{id}", u.GetUser)
	app.Mount("/users", resources.Users{}.Routes())

	// start server
	// defer http.ListenAndServe(":3001", app)
	fmt.Printf("[%s]: Server running...", appname)
	err := http.ListenAndServe(host+port, app)
	if err != nil {
		log.Fatal("Serving: ", err)
	}
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s
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
