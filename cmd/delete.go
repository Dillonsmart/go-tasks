package cmd

import "github.com/Dillonsmart/go-tasks/db"

func Delete(taskID int) {
	query := "DELETE FROM tasks WHERE id = ?"
	_, err := db.DB.Exec(query, taskID)

	if err != nil {
		println("Error deleting task:", err)
		return
	}

	println("Task deleted:", taskID)
}
