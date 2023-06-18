package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/marcoscoutinhodev/url_shortener_api/docs"
	"github.com/marcoscoutinhodev/url_shortener_api/external/handler"
	"github.com/marcoscoutinhodev/url_shortener_api/external/middlewares"
	httpSwager "github.com/swaggo/http-swagger"
)

// @Title						URL SHORTENER API
// @version					0.1
// @description			api for url shortener application
// @termsOfServices	https://swagger.io/terms/

// @contact.name	Marcos Coutinho
// @contact.url		https://linkedin.com/in/marcoscoutinhodev
// @contact.email marcoscoutinhodev@outlook.com

// @license.name	The MIT License (MIT)
// @license.url		https://mit-license.org/

// @host												localhost:4001
// @BashPath										/
// @securityDefinitions.apiKey  ApiKeyAuth
// @in													header
// @name												x-access-token
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
		r.Post("/", middlewares.AuthenticationMiddleware(handler.CreateShortURL).(http.HandlerFunc))
		r.Get("/{shortURL}", handler.GetOriginalURL)
		r.Patch("/report/{urlID}", middlewares.AuthenticationMiddleware(handler.ReportURL).(http.HandlerFunc))
		r.Patch("/active/{urlID}", middlewares.AuthenticationMiddleware(handler.ActiveURL).(http.HandlerFunc))
	})

	r.Get("/docs/*", httpSwager.Handler(httpSwager.URL(
		fmt.Sprintf("http://localhost:%s/docs/doc.json", os.Getenv("SERVER_PORT")),
	)))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), r); err != nil {
		panic(err)
	}
}
