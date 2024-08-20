package handlers

import (
	"github.com/scottmckendry/mnemstart/auth"
	"github.com/scottmckendry/mnemstart/data"
)

type Handler struct {
	store *data.Storage
	auth  *auth.AuthService
}

func New(store *data.Storage, auth *auth.AuthService) *Handler {
	return &Handler{
		store: store,
		auth:  auth,
	}
}
