package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (c *Config) server() {
	mainRouter := chi.NewRouter()

	// Use the good base middleware stack from chi recommendation
	mainRouter.Use(middleware.RequestID)
	mainRouter.Use(middleware.RealIP)
	mainRouter.Use(middleware.Logger)
	mainRouter.Use(middleware.Recoverer)
	mainRouter.Use(c.serveFilesHandler())

	mainRouter.Get("/", c.homeHandler)
	// Lets add css directory to the server
	serv := &http.Server{
		Addr:        ":" + c.PORT,
		Handler:     mainRouter,
		ReadTimeout: time.Minute,
	}
	log.Println("Starting the server on port", c.PORT)
	log.Fatal(serv.ListenAndServe())
}
