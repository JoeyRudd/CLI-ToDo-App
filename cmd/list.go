package cmd

import (
	"fmt"
	"github.com/JoeyRudd/CLI-ToDo-App/internal"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

var showAll bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long: `The list command retrieves and displays all 
			tasks from the task list.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := internal.GetAllTasksFromDB(internal.DB)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading tasks:", err)
			return
		}
		if len(tasks) == 0 {
			fmt.Fprintln(os.Stderr, "No tasks found.")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tDescription\tCreated\tCompleted")
		for _, task := range tasks {
			// Check if the task should be shown based on the --all flag
			if showAll || !task.Completed {
				completed := "No"
				if task.Completed {
					completed = "Yes"
				}
				fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.ID, task.Description, internal.FormatTimeAsAgo(task.Created), completed)
			}
		}
		w.Flush()
	},
}

func init() {
	// added a show all command
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all tasks including completed ones")
	rootCmd.AddCommand(listCmd)
}
