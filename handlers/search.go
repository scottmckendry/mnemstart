package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/scottmckendry/mnemstart/views"
)

func (h *Handler) HandleSearchSuggest(w http.ResponseWriter, r *http.Request) {
	engine := r.FormValue("search_engine")
	query := r.FormValue("q")
	query = url.QueryEscape(query)

	body, err := getGoogleSuggestions(query)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error getting suggestions: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	var suggestions []any
	err = json.Unmarshal(body, &suggestions)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error unmarshalling suggestions: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	parsedSuggestions := parseSuggestions(suggestions[1].([]any))
	views.Suggestions(parsedSuggestions, engine).Render(r.Context(), w)
}

func getGoogleSuggestions(query string) ([]byte, error) {
	resp, err := http.Get(
		"https://suggestqueries.google.com/complete/search?client=firefox&q=" + query,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func parseSuggestions(suggestions []any) []string {
	var parsedSuggestions []string
	for _, suggestion := range suggestions {
		parsedSuggestions = append(parsedSuggestions, suggestion.(string))
	}

	return parsedSuggestions
}
