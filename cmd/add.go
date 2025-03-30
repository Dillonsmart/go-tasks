package cmd

import (
	"fmt"
	"github.com/Dillonsmart/go-tasks/db"
)

const TaskNameMaxLength = 255

func Add(taskName string) {
	if len(taskName) == 0 {
		fmt.Println("Please provide a task to add.")
		return
	}

	if len(taskName) > TaskNameMaxLength {
		fmt.Println("Task name is too long. Please limit it to 255 characters.")
		return
	}

	if TaskExists(taskName) {
		fmt.Printf("Task \"%s\" already exists.", taskName)
		return
	}

	addTask(taskName)
}

func addTask(taskName string) {
	query := "INSERT INTO tasks (name) VALUES (?)"
	_, err := db.DB.Exec(query, taskName)

	if err != nil {
		fmt.Println("Error adding task:", err)
		return
	}

	fmt.Println("Task added:", taskName)
}
