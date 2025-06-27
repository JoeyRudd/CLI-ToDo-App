package cmd

import (
	"fmt"
	"github.com/JoeyRudd/CLI-ToDo-App/internal"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
				Completed:   false,
			}

			if err := internal.AppendTaskToCSV(task, filename); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		} else {
			fmt.Println("Please provide a task description.")
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
