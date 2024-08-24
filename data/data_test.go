package data

import (
	"strconv"
	"testing"

	"github.com/markbates/goth"
)

var db, _ = NewLibSqlDatabase("file:test.db")
var store = NewStore(db)
var userEmail = "john@example.com"

func TestCreateOrUpdateUser(t *testing.T) {
	user := goth.User{
		Provider: "github",
		Email:    userEmail,
		Name:     "John Doe",
		UserID:   "1234",
	}

	err := store.CreateOrUpdateUser(user)
	if err != nil {
		t.Errorf("Error creating or updating user: %v", err)
	}

	mnemstart_user := User{
		Email: userEmail,
	}

	err = getUser(db, &mnemstart_user)
	if err != nil {
		t.Errorf("Error getting user: %v", err)
	}

	if mnemstart_user.GithubID != "1234" {
		t.Errorf("Expected GithubID to be 1234, got %s", mnemstart_user.GithubID)
	}

	user2 := goth.User{
		Provider: "discord",
		Email:    mnemstart_user.Email,
		UserID:   "5678",
	}

	err = store.CreateOrUpdateUser(user2)
	if err != nil {
		t.Errorf("Error creating or updating user: %v", err)
	}

	err = getUser(db, &mnemstart_user)
	if err != nil {
		t.Errorf("Error getting user: %v", err)
	}

	if mnemstart_user.DiscordID != "5678" {
		t.Errorf("Expected DiscordID to be 5678, got %s", mnemstart_user.DiscordID)
	}
}

func TestGetUser(t *testing.T) {
	user := User{
		Email: userEmail,
	}

	err := getUser(db, &user)

	if err != nil {
		t.Errorf("Error getting user: %v", err)
	}

	if user.GithubID != "1234" {
		t.Errorf("Expected GithubID to be 1234, got %s", user.GithubID)
	}

	if user.DiscordID != "5678" {
		t.Errorf("Expected DiscordID to be 5678, got %s", user.DiscordID)
	}
}

func TestAddMapping(t *testing.T) {
	_, err := db.Exec("DELETE FROM mappings")

	mapping := Mapping{
		Keymap: "gh",
		MapsTo: "https://github.com",
	}

	err = store.AddMapping(userEmail, mapping.Keymap, mapping.MapsTo)
	if err != nil {
		t.Errorf("Error adding mapping: %v", err)
	}

	mappings := store.GetMappings(userEmail)
	if err != nil {
		t.Errorf("Error getting mappings: %v", err)
	}

	if len(mappings) != 1 {
		t.Errorf("Expected 1 mapping, got %d", len(mappings))
	}

	if mappings[0].Keymap != mapping.Keymap {
		t.Errorf("Expected keymap to be %s, got %s", mapping.Keymap, mappings[0].Keymap)
	}

	if mappings[0].MapsTo != mapping.MapsTo {
		t.Errorf("Expected maps_to to be %s, got %s", mapping.MapsTo, mappings[0].MapsTo)
	}
}

func TestGetMappings(t *testing.T) {
	mappings := store.GetMappings(userEmail)

	if len(mappings) != 1 {
		t.Errorf("Expected 1 mapping, got %d", len(mappings))
	}

	if mappings[0].Keymap != "gh" {
		t.Errorf("Expected keymap to be gh, got %s", mappings[0].Keymap)
	}

	if mappings[0].MapsTo != "https://github.com" {
		t.Errorf("Expected maps_to to be https://github.com, got %s", mappings[0].MapsTo)
	}
}

func TestUpdateMapping(t *testing.T) {
	mappings := store.GetMappings(userEmail)
	if len(mappings) != 1 {
		t.Errorf("Expected 1 mapping, got %d", len(mappings))
	}

	mapping := mappings[0]
	mapping.MapsTo = "https://gitlab.com"

	err := store.UpdateMapping(strconv.Itoa(mapping.ID), userEmail, mapping.Keymap, mapping.MapsTo)
	if err != nil {
		t.Errorf("Error updating mapping: %v", err)
	}

	mappings = store.GetMappings(userEmail)
	if len(mappings) != 1 {
		t.Errorf("Expected 1 mapping, got %d", len(mappings))
	}

	if mappings[0].MapsTo != "https://gitlab.com" {
		t.Errorf("Expected maps_to to be https://gitlab.com, got %s", mappings[0].MapsTo)
	}
}

func TestDeleteMapping(t *testing.T) {
	mappings := store.GetMappings(userEmail)
	if len(mappings) != 1 {
		t.Errorf("Expected 1 mapping, got %d", len(mappings))
	}

	mapping := mappings[0]

	err := store.DeleteMapping(strconv.Itoa(mapping.ID), userEmail)
	if err != nil {
		t.Errorf("Error deleting mapping: %v", err)
	}

	mappings = store.GetMappings(userEmail)
	if len(mappings) != 0 {
		t.Errorf("Expected 0 mappings, got %d", len(mappings))
	}
}

func TestUpdateUserSettings(t *testing.T) {
	user := User{
		Email: userEmail,
	}

	err := getUser(db, &user)
	if err != nil {
		t.Errorf("Error getting user: %v", err)
	}

	settings := UserSettings{
		SearchEngine: "test_search_engine",
		LeaderKey:    "test_leader_key",
	}

	err = store.UpdateUserSettings(user.Email, &settings)
	if err != nil {
		t.Errorf("Error updating user settings: %v", err)
	}

	err = getUser(db, &user)
	if err != nil {
		t.Errorf("Error getting user: %v", err)
	}

	returnedSettings := store.GetUserSettings(user.Email)
	if returnedSettings.SearchEngine != settings.SearchEngine {
		t.Errorf(
			"Expected search engine to be %s, got %s",
			settings.SearchEngine,
			returnedSettings.SearchEngine,
		)
	}

	if returnedSettings.LeaderKey != settings.LeaderKey {
		t.Errorf(
			"Expected leader key to be %s, got %s",
			settings.LeaderKey,
			returnedSettings.LeaderKey,
		)
	}
}

func TestGetUserSettings(t *testing.T) {
	user := User{
		Email: userEmail,
	}

	settings := store.GetUserSettings(user.Email)
	if settings.SearchEngine != "test_search_engine" {
		t.Errorf("Expected search engine to be test_search_engine, got %s", settings.SearchEngine)
	}

	if settings.LeaderKey != "test_leader_key" {
		t.Errorf("Expected leader key to be test_leader_key, got %s", settings.LeaderKey)
	}
}
