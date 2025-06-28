package main

import (
	"github.com/JoeyRudd/CLI-ToDo-App/cmd"
	"github.com/JoeyRudd/CLI-ToDo-App/internal"
	"log"
)

func main() {
	// Initialize the database connection
	var err error
	internal.DB, err = internal.InitDB("tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	// Ensure the database connection is closed when the application exits
	defer internal.DB.Close()
	cmd.Execute()
}
