package main

import (
	"testing"

	"github.com/scottmckendry/mnemstart/data"
)

func TestDbInit(t *testing.T) {
	db, _ := data.NewLibSqlDatabase("file:test.db")
	initStorage(db)

	err := db.Ping()
	if err != nil {
		t.Errorf("Error pinging database: %v", err)
	}
}
