package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/scottmckendry/mnemstart/auth"
	"github.com/scottmckendry/mnemstart/config"
	"github.com/scottmckendry/mnemstart/data"
	"github.com/scottmckendry/mnemstart/handlers"
)

func main() {
	db, err := data.NewLibSqlDatabase(config.Envs.DatabaseURL)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	store := data.NewStore(db)

	initStorage(db)

	sessionStore := auth.NewCookieStore(auth.SessionOptions{
		CookiesKey: config.Envs.CookiesAuthSecret,
		MaxAge:     config.Envs.CookiesAuthAgeInSeconds,
		HttpOnly:   config.Envs.CookiesAuthIsHttpOnly,
		Secure:     config.Envs.CookiesAuthIsSecure,
	})
	authService := auth.NewAuthService(sessionStore)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	handler := handlers.New(store, authService)

	r.Group(func(r chi.Router) {
		// Require authentication for all routes in this group
		r.Use(auth.RequireAuth(authService))

		// app routes
		r.Get("/", handler.HandleRoot)
		r.Get("/settings", handler.HandleSettings)
		r.Put("/update-settings", handler.HandleSettingsUpdate)
		r.Get("/mappings", handler.HandleMappings)
		r.Get("/mappings/{id}", handler.HandleMapping)
		r.Get("/mappings/new", handler.HandleMappingNew)
		r.Post("/mappings/add", handler.HandleMappingAdd)
		r.Get("/mappings/edit/{id}", handler.HandleMappingEdit)
		r.Put("/mappings/update/{id}", handler.HandleMappingUpdate)
		r.Delete("/mappings/delete/{id}", handler.HandleMappingDelete)
		r.Post("/search/suggest", handler.HandleSearchSuggest)
		r.Get("/help", handler.HandleHelp)
	})

	// auth
	r.Get("/auth/{provider}", handler.HandleProviderLogin)
	r.Get("/auth/{provider}/callback", handler.HandleAuthCallbackFunction)
	r.Get("/auth/{provider}/logout", handler.HandleLogout)
	r.Get("/login", handler.HandleLogin)

	// static content
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Printf("Server: Listening on %s:%s", config.Envs.PublicHost, config.Envs.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Envs.Port), r))
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("DB: Connected!")
}
