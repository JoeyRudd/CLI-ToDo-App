package cmd

import (
	"fmt"
	"github.com/JoeyRudd/CLI-ToDo-App/internal"
	"github.com/spf13/cobra"
	"strconv"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task",
	Long:  `The complete command marks a task as completed.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check if a task ID is provided
		if len(args) < 1 {
			fmt.Println("Please provide a task ID to complete.")
			return
		}

		// convert the first argument to an integer task ID
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID:", args[0])
			return
		}

		filename := "tasks.csv"
		// Update the task in the CSV file
		err = internal.UpdateTaskInCSV(taskID, filename)
		if err != nil {
			fmt.Println("Error completing task:", err)
			return
		}
		fmt.Printf("Task %s marked as completed.\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
