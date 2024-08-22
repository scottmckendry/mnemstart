package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

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

func (h *Handler) HandleMappings(w http.ResponseWriter, r *http.Request) {
	user, err := h.auth.GetSessionUser(r)
	if err != nil {
		log.Println(err)
		return
	}

	mappings := h.store.GetMappings(user.Email)
	views.Mappings(user, mappings).Render(r.Context(), w)
}

func (h *Handler) HandleMapping(w http.ResponseWriter, r *http.Request) {
	user, err := h.auth.GetSessionUser(r)
	if err != nil {
		log.Println(err)
		return
	}

	mappingID := chi.URLParam(r, "id")
	mapping := h.store.GetMapping(mappingID, user.Email)
	views.MappingRow(mapping).Render(r.Context(), w)
}

func (h *Handler) HandleMappingNew(w http.ResponseWriter, r *http.Request) {
	views.NewMapping().Render(r.Context(), w)
}

func (h *Handler) HandleMappingAdd(w http.ResponseWriter, r *http.Request) {
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

	keymap := r.FormValue("keymap")
	mapsTo := r.FormValue("mapsto")
	h.store.AddMapping(user.Email, keymap, mapsTo)
	mappings := h.store.GetMappings(user.Email)
	views.Mappings(user, mappings).Render(r.Context(), w)
}

func (h *Handler) HandleMappingDelete(w http.ResponseWriter, r *http.Request) {
	user, err := h.auth.GetSessionUser(r)
	if err != nil {
		log.Println(err)
		return
	}

	mappingID := chi.URLParam(r, "id")
	h.store.DeleteMapping(mappingID, user.Email)
	mappings := h.store.GetMappings(user.Email)
	views.Mappings(user, mappings).Render(r.Context(), w)
}

func (h *Handler) HandleMappingEdit(w http.ResponseWriter, r *http.Request) {
	user, err := h.auth.GetSessionUser(r)
	if err != nil {
		log.Println(err)
		return
	}

	mappingID := chi.URLParam(r, "id")
	mapping := h.store.GetMapping(mappingID, user.Email)
	views.EditMapping(user, mapping).Render(r.Context(), w)
}

func (h *Handler) HandleMappingUpdate(w http.ResponseWriter, r *http.Request) {
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

	mappingID := chi.URLParam(r, "id")
	keymap := r.FormValue("keymap")
	mapsTo := r.FormValue("mapsto")
	h.store.UpdateMapping(mappingID, user.Email, keymap, mapsTo)
	mappings := h.store.GetMappings(user.Email)
	views.Mappings(user, mappings).Render(r.Context(), w)
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
	h.store.UpdateUserSettings(user.Email, userSettings)

	mappings := h.store.GetMappings(user.Email)

	w.Header().Set("HX-Refresh", "true")
	views.Home(user, userSettings, mappings).Render(r.Context(), w)
}
