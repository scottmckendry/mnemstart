package handlers

import (
	"log"
	"net/http"

	"github.com/scottmckendry/mnemstart/views"
)

func (h *Handler) HandleRoot(w http.ResponseWriter, r *http.Request) {
	user, err := h.auth.GetSessionUser(r)
	if err != nil {
		log.Println(err)
		return
	}

	userSettings := h.store.GetUserSettings(user.Email)

	views.Home(user, userSettings).Render(r.Context(), w)
}
