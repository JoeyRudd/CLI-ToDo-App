package cmd

import (
	"fmt"
	"github.com/JoeyRudd/CLI-ToDo-App/internal"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long: `The list command retrieves and displays all 
			tasks from the task list.`,
	Run: func(cmd *cobra.Command, args []string) {
		filename := "tasks.csv"
		tasks, err := internal.ReadTasksFromCSV(filename)
		if err != nil {
			fmt.Println("Error reading tasks:", err)
			return
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		fmt.Println("Tasks:")
		for _, task := range tasks {
			fmt.Printf("ID: %d, Description: %s, Completed: %t\n",
				task.ID, task.Description, task.Completed)
		}
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
