package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"

	"github.com/scottmckendry/mnemstart/views"
)

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	views.Login().Render(r.Context(), w)
}

func (h *Handler) HandleProviderLogin(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	// try to get the user without re-authenticating
	if u, err := gothic.CompleteUserAuth(w, r); err == nil {
		slog.Info("user already authenticated", "user", u)
		views.Login().Render(r.Context(), w)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func (h *Handler) HandleAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		slog.Error("failed to complete user authentication", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.auth.StoreUserSession(w, r, user)
	if err != nil {
		slog.Error("failed to store user session", "error", err)
		return
	}

	err = h.store.CreateOrUpdateUser(user)
	if err != nil {
		slog.Error("failed to create or update user", "error", err)
		return
	}

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	slog.Info("logging out user")

	err := gothic.Logout(w, r)
	if err != nil {
		slog.Error("failed to logout user", "error", err)
		return
	}

	h.auth.RemoveUserSession(w, r)

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
