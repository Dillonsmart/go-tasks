package cmd

import (
	"fmt"
	"github.com/Dillonsmart/go-tasks/db"
)

func List() {
	query := "SELECT id, name FROM tasks"
	rows, err := db.DB.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Printf("(%d) %s\n", id, name)
	}
}
