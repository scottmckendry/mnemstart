package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/scottmckendry/mnemstart/auth"
	"github.com/scottmckendry/mnemstart/config"
	"github.com/scottmckendry/mnemstart/data"
	"github.com/scottmckendry/mnemstart/handlers"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	db, err := data.NewLibSqlDatabase(config.Envs.DatabaseURL)
	if err != nil {
		slog.Error("failed to open database", "error", err)
		os.Exit(1)
	}

	store := data.NewStore(db)

	initStorage(db)

	sessionStore, err := auth.NewFileStore(auth.SessionOptions{
		StorePath:  "./sessions",
		CookiesKey: config.Envs.CookiesAuthSecret,
		MaxAge:     config.Envs.CookiesAuthAgeInSeconds,
		HttpOnly:   config.Envs.CookiesAuthIsHttpOnly,
		Secure:     config.Envs.CookiesAuthIsSecure,
	})
	if err != nil {
		slog.Error("failed to create session store", "error", err)
		os.Exit(1)
	}
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

	slog.Info("server starting",
		"host", config.Envs.PublicHost,
		"port", config.Envs.Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Envs.Port), r); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		slog.Error("failed to ping database", "error", err)
		os.Exit(1)
	}

	slog.Info("database connected")
}
