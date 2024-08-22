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
	r.HandleFunc("/", auth.RequireAuth(handler.HandleRoot, authService))
	r.HandleFunc("/settings", auth.RequireAuth(handler.HandleSettings, authService))
	r.HandleFunc("/update-settings", auth.RequireAuth(handler.HandleSettingsUpdate, authService))
	r.HandleFunc("/mappings", auth.RequireAuth(handler.HandleMappings, authService))
	r.HandleFunc("/mappings/{id}", auth.RequireAuth(handler.HandleMapping, authService))
	r.HandleFunc("/mappings/new", auth.RequireAuth(handler.HandleMappingNew, authService))
	r.HandleFunc("/mappings/add", auth.RequireAuth(handler.HandleMappingAdd, authService))
	r.HandleFunc("/mappings/edit/{id}", auth.RequireAuth(handler.HandleMappingEdit, authService))
	r.HandleFunc(
		"/mappings/update/{id}",
		auth.RequireAuth(handler.HandleMappingUpdate, authService),
	)
	r.HandleFunc(
		"/mappings/delete/{id}",
		auth.RequireAuth(handler.HandleMappingDelete, authService),
	)

	// auth
	r.HandleFunc("/auth/{provider}", handler.HandleProviderLogin)
	r.HandleFunc("/auth/{provider}/callback", handler.HandleAuthCallbackFunction)
	r.HandleFunc("/auth/{provider}/logout", handler.HandleLogout)
	r.HandleFunc("/login", handler.HandleLogin)

	// static content
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Printf("Server: Listening on %s:%s", config.Envs.PublicHost, config.Envs.Port)
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%s", config.Envs.Port), r),
	)
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("DB: Connected!")
}
