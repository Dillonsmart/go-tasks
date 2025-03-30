package main

import (
	"fmt"
	"os"

	"github.com/Dillonsmart/go-tasks/cmd"
)

func main() {
	args := os.Args[1:]

	action := args[0] // The first arguement is always the action

	switch action {
	case "add":
		if len(args) < 2 {
			fmt.Println("Please provide a task to add.")
			return
		}
		cmd.Add(args[1])
		break
	}
}
