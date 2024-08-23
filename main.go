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

	handler := handlers.New(store, authService)

	// app routes
	r.Get("/", auth.RequireAuth(handler.HandleRoot, authService))
	r.Get("/settings", auth.RequireAuth(handler.HandleSettings, authService))
	r.Put("/update-settings", auth.RequireAuth(handler.HandleSettingsUpdate, authService))
	r.Get("/mappings", auth.RequireAuth(handler.HandleMappings, authService))
	r.Get("/mappings/{id}", auth.RequireAuth(handler.HandleMapping, authService))
	r.Get("/mappings/new", auth.RequireAuth(handler.HandleMappingNew, authService))
	r.Post("/mappings/add", auth.RequireAuth(handler.HandleMappingAdd, authService))
	r.Get("/mappings/edit/{id}", auth.RequireAuth(handler.HandleMappingEdit, authService))
	r.Put("/mappings/update/{id}", auth.RequireAuth(handler.HandleMappingUpdate, authService))
	r.Delete("/mappings/delete/{id}", auth.RequireAuth(handler.HandleMappingDelete, authService))

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
