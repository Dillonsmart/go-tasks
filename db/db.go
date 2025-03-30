package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func Init() {
	// Initialize the database connection
	fmt.Println("Initializing database")

	dbPath := "./db/tasks.db"
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	fmt.Println("Successfully opened database")
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
	} else {
		fmt.Println("Successfully created tasks table")
	}
}
