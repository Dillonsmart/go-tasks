package main

import (
	"fmt"
	"os"

	"github.com/Dillonsmart/go-tasks/cmd"
	"github.com/Dillonsmart/go-tasks/db"
)

func main() {
	db.Init()

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please provide an action.")
		return
	}

	action := args[0] // The first argument is always the action to perform

	switch action {
	case "add":
		if len(args) < 2 {
			fmt.Println("Please provide a task to add.")
			return
		}
		cmd.Add(args[1])
		break
	case "list":
		cmd.List()
		break
	}
}
