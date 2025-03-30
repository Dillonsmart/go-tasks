package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func Init() {
	dbPath := "./db/tasks.db"
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	DB = db

	createTasksTable()
}

func createTasksTable() {
	query := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255) NOT NULL
		);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create tasks table: %v", err)
	}
}
