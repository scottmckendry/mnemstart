package data

import (
	"testing"
)

func TestDbInit(t *testing.T) {
	db, _ := NewLibSqlDatabase("file:test.db")
	err := db.Ping()
	if err != nil {
		t.Errorf("Error pinging database: %v", err)
	}
}

func TestGenerateSchema(t *testing.T) {
	db, _ := NewLibSqlDatabase("file:test.db")
	err := generateSchema(db)
	if err != nil {
		t.Errorf("Error generating schema: %v", err)
	}

	tables := []string{"users", "mappings", "user_settings"}
	for _, table := range tables {
		t.Run(table, func(t *testing.T) {
			rows, err := db.Query(
				"SELECT name FROM sqlite_master WHERE type='table' AND name=?",
				table,
			)
			if err != nil {
				t.Errorf("Error querying for table %s: %v", table, err)
			}
			defer rows.Close()

			if !rows.Next() {
				t.Errorf("Table %s not found", table)
			}
		})
	}
}
