package cmd

import (
	"github.com/JoeyRudd/CLI-ToDo-App/internal"
	"github.com/spf13/cobra"
	"time"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new task",
	Long:  `add command allows you to create a new task with a description.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			description := args[0]
			// Here you would typically call a function to add the task to your database
			task := internal.Task{
				Description: description,
				Created:     time.Now(),
				Completed:   false,
			}
			err := internal.AddTaskToDB(internal.DB, task)
			if err != nil {
				cmd.Println("Error adding task:", err)
			} else {
				cmd.Println("Task added successfully:", task.Description)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
