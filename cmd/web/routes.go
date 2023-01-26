package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/suisanT1/bookings/pkg/config"
	"github.com/suisanT1/bookings/pkg/handlers"
)

// Creating routes using 3rd party package pat
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Add Middleware
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/main", handlers.Repo.Main)

	return mux
}
