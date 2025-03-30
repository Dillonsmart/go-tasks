package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Dillonsmart/go-tasks/db"
)

func TaskExists(taskName string) bool {
	query := "SELECT 1 FROM tasks WHERE name = ? LIMIT 1"
	row := db.DB.QueryRow(query, taskName)
	var exists int

	err := row.Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}
		fmt.Println("Error checking task existence:", err)
		return false
	}

	return true
}
