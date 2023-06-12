package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/marcoscoutinhodev/url_shortener_api/external/handler"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/user", func(r chi.Router) {
		r.Post("/signup", handler.CreateUser)
		r.Post("/signin", handler.AuthenticateUser)
	})

	r.Route("/url", func(r chi.Router) {
		r.Post("/", handler.CreateShortURL)
	})

	if err := http.ListenAndServe(os.Getenv("SERVER_PORT"), r); err != nil {
		panic(err)
	}
}
