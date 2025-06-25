package main

import (
	"fmt"
	"github.com/JoeyRudd/CLI-ToDo-App/internal"
)

func main() {
	fmt.Printf("Welcome to the CLI To-Do App!\n")

	// testing
	task := internal.Task{ID: 1, Description: "Learn Go", Completed: false}
	err := internal.AppendTaskToCSV(task, "tasks.csv")
	if err != nil {
		fmt.Printf("Error appending task to CSV: %v\n", err)
	} else {
		fmt.Printf("Task added successfully: %v\n", task)
	}
}
