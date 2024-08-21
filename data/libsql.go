package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

func NewLibSqlDatabase(url string) (*sql.DB, error) {
	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = generateSchema(db)
	if err != nil {
		return nil, fmt.Errorf("Error generating schema: %v", err)
	}

	return db, nil
}

func generateSchema(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY,
            name TEXT,
            email TEXT,
            discord_id TEXT,
            github_id TEXT
        );
    `)
	if err != nil {
		return fmt.Errorf("Error creating users table: %v", err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS mappings (
            id INTEGER PRIMARY KEY,
            user_id INTEGER,
            keymap TEXT,
            maps_to TEXT
        );
    `)
	if err != nil {
		return fmt.Errorf("Error creating mappings table: %v", err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS user_settings (
            id INTEGER PRIMARY KEY,
            user_id INTEGER,
            setting_key TEXT,
            setting_value TEXT
        );
    `)
	if err != nil {
		return fmt.Errorf("Error creating user_settings table: %v", err)
	}

	return nil
}
