package cmd

import (
	"fmt"
	"github.com/JoeyRudd/CLI-ToDo-App/internal"
	"github.com/spf13/cobra"
	"os"
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
			filename := "tasks.csv"
			// Create a new task with the provided description
			tasks, err := internal.ReadTasksFromCSV(filename)
			nextID := 1
			if err == nil && len(tasks) > 0 {
				nextID = tasks[len(tasks)-1].ID + 1
			}

			task := internal.Task{
				ID:          nextID,
				Description: description,
				Created:     time.Now(),
				Completed:   false,
			}

			if err := internal.AppendTaskToCSV(task, filename); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

		} else {
			fmt.Fprintln(os.Stderr, "Please provide a task description.")
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
