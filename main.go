package main

import (
	"fmt"
	"github.com/JoeyRudd/CLI-ToDo-App/internal"
	"os"
)

func main() {

	os.Remove("tasks.csv") // Remove the file if it exists to start fresh

	fmt.Printf("Welcome to the CLI To-Do App!\n")

	// testing for task1
	task1 := internal.Task{ID: 1, Description: "Learn Go", Completed: false}
	err := internal.AppendTaskToCSV(task1, "tasks.csv")
	if err != nil {
		fmt.Printf("Error appending task to CSV: %v\n", err)
	} else {
		fmt.Printf("Task added successfully: %v\n", task1)
	}

	// testing for task2
	task2 := internal.Task{ID: 2, Description: "Code Go", Completed: false}
	err = internal.AppendTaskToCSV(task2, "tasks.csv")
	if err != nil {
		fmt.Printf("Error appending task to CSV: %v\n", err)
	} else {
		fmt.Printf("Task added successfully: %v\n", task2)
	}

	// testing for task3
	task3 := internal.Task{ID: 3, Description: "Master Go", Completed: false}
	err = internal.AppendTaskToCSV(task3, "tasks.csv")
	if err != nil {
		fmt.Printf("Error appending task to CSV: %v\n", err)
	} else {
		fmt.Printf("Task added successfully: %v\n", task3)
	}

	// testing reading tasks from CSV
	tasks, err := internal.ReadTasksFromCSV("tasks.csv")
	if err != nil {
		fmt.Printf("Error reading tasks from CSV: %v\n", err)
	} else {
		fmt.Println("Tasks read from CSV: ")
		for _, task := range tasks {
			fmt.Printf("%v\n", task)
		}
	}

	// testing updating a task in CSV
	internal.UpdateTaskInCSV(task1, "tasks.csv")
	tasks, err = internal.ReadTasksFromCSV("tasks.csv")
	fmt.Printf("Testing code after updating task1:\n")
	for _, task := range tasks {
		fmt.Printf("%v\n", task)
	}

}
