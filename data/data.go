package data

import (
	"database/sql"

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

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateUser(gothUser goth.User) error {
	user := buildUserFromGothUser(gothUser)
	err := getUser(s.db, user)
	if err != nil {
		err = createUser(s.db, user)
		if err != nil {
			return err
		}
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
