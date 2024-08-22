package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/markbates/goth"
)

type Storage struct {
	db *sql.DB
}

type User struct {
	ID        int
	Name      string
	Email     string
	DiscordID string
	GithubID  string
}

type Mapping struct {
	ID     int
	Keymap string
	MapsTo string
}

type UserSettings struct {
	SearchEngine string
	LeaderKey    string
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateOrUpdateUser(gothUser goth.User) error {
	user := buildUserFromGothUser(gothUser)
	err := getUser(s.db, user)
	if err != nil {
		err = createUser(s.db, user)
		if err != nil {
			return fmt.Errorf("Error creating user: %v", err)
		}
	}

	err = updateUser(s.db, user)
	if err != nil {
		return fmt.Errorf("Error updating user: %v", err)
	}

	return nil
}

func getUser(db *sql.DB, user *User) error {
	row := db.QueryRow("SELECT * FROM users WHERE email = ?", user.Email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.DiscordID, &user.GithubID)
	if err != nil {
		return err
	}

	return nil
}

func createUser(db *sql.DB, user *User) error {
	_, err := db.Exec(
		"INSERT INTO users (name, email, discord_id, github_id) VALUES (?, ?, ?, ?)",
		user.Name,
		user.Email,
		user.DiscordID,
		user.GithubID,
	)
	if err != nil {
		return err
	}

	return nil
}

func updateUser(db *sql.DB, user *User) error {
	_, err := db.Exec(`
        UPDATE users
        SET
            name = CASE
                WHEN name IS NULL THEN ?
                ELSE name
            END,
            discord_id = ?,
            github_id = ?
        WHERE email = ?
    `,
		user.Name,
		user.DiscordID,
		user.GithubID,
		user.Email,
	)
	if err != nil {
		return err
	}

	return nil
}

func buildUserFromGothUser(gothUser goth.User) *User {
	user := &User{
		Name:  gothUser.Name,
		Email: gothUser.Email,
	}

	switch gothUser.Provider {
	case "discord":
		user.DiscordID = gothUser.UserID
	case "github":
		user.GithubID = gothUser.UserID
	}

	return user
}

func (s *Storage) GetMappings(email string) []Mapping {
	mappings := []Mapping{}
	rows, err := s.db.Query(
		`SELECT mappings.id, keymap, maps_to
            FROM mappings
            INNER JOIN users
            ON mappings.user_id = users.id
            WHERE users.email = ?`,
		email,
	)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		mapping := Mapping{}
		err = rows.Scan(&mapping.ID, &mapping.Keymap, &mapping.MapsTo)
		if err != nil {
			return nil
		}

		mappings = append(mappings, mapping)
	}

	return mappings
}

func (s *Storage) GetUserSettings(email string) *UserSettings {
	settings := &UserSettings{}
	rows, err := s.db.Query(
		`SELECT setting_key, setting_value
            FROM user_settings
            INNER JOIN users
            ON user_settings.user_id = users.id
            WHERE users.email = ?`,
		email,
	)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var key, value string
		err = rows.Scan(&key, &value)
		if err != nil {
			return nil
		}

		switch key {
		case "SearchEngine":
			settings.SearchEngine = value
		case "LeaderKey":
			settings.LeaderKey = value
		}
	}

	return settings
}

func (s *Storage) UpdateUserSettings(email string, settings *UserSettings) error {
	settingsMap := map[string]interface{}{
		"SearchEngine": settings.SearchEngine,
		"LeaderKey":    settings.LeaderKey,
	}

	for settingKey, settingValue := range settingsMap {
		_, err := s.db.Exec(
			`INSERT OR REPLACE INTO user_settings (user_id, setting_key, setting_value)
                VALUES (
                    (SELECT id FROM users WHERE email = ?),
                    ?,
                    ?
                )`,
			email,
			settingKey,
			settingValue,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
