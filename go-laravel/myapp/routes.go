package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes

	// add routes here
	a.App.Routes.Get("/", a.Handlers.Home)

	a.App.Routes.Get("/jet", func(w http.ResponseWriter, r *http.Request) {
		log.Println("/jet called")
		_ = a.App.Render.JetPage(w, r, "test-jet", nil, nil)
	})

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
