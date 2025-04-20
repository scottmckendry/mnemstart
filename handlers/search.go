package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/scottmckendry/mnemstart/views"
)

func (h *Handler) HandleSearchSuggest(w http.ResponseWriter, r *http.Request) {
	engine := r.FormValue("search_engine")
	query := r.FormValue("q")
	query = url.QueryEscape(query)

	slog.Info("getting search suggestions",
		"engine", engine,
		"query", query)

	body, err := getGoogleSuggestions(query)
	if err != nil {
		slog.Error("failed to get suggestions from Google API",
			"error", err,
			"query", query)
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
		slog.Error("failed to unmarshal suggestions",
			"error", err,
			"body", string(body))
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
	url := "https://suggestqueries.google.com/complete/search?client=firefox&q=" + query
	slog.Debug("calling Google suggestions API", "url", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call Google API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
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
