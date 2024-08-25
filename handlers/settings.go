package handlers

import (
	"log"
	"net/http"

	"github.com/scottmckendry/mnemstart/views"
)

func (h *Handler) HandleSettings(w http.ResponseWriter, r *http.Request) {
	user, err := h.auth.GetSessionUser(r)
	if err != nil {
		log.Println(err)
		return
	}

	userSettings := h.store.GetUserSettings(user.Email)
	views.Settings(user, userSettings).Render(r.Context(), w)
}

func (h *Handler) HandleSettingsUpdate(w http.ResponseWriter, r *http.Request) {
	user, err := h.auth.GetSessionUser(r)
	if err != nil {
		log.Println(err)
		return
	}

	err = r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	userSettings := h.store.GetUserSettings(user.Email)
	userSettings.LeaderKey = r.FormValue("leaderKey")
	userSettings.SearchEngine = r.FormValue("searchEngine")

	log.Println(userSettings)
	h.store.UpdateUserSettings(user.Email, userSettings)

	mappings := h.store.GetMappings(user.Email)

	w.Header().Set("HX-Refresh", "true")
	views.Home(user, userSettings, mappings).Render(r.Context(), w)
}
