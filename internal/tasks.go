package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
	"log"
	"os"
	"time"
)

type Task struct {
	ID          int
	Description string
	Created     time.Time
	Completed   bool
}

// CloseFile Helper function to close a file and log an error if it occurs			fmt.Fprintf(w, "%d\t%s\t%s\t%t\n", task.ID, task.Description, task.Created, task.Completed)
func CloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Printf("error closing file: %v", err)
	}
}

// AddTaskToDB adds a new task to the database
func AddTaskToDB(db *sql.DB, task Task) error {
	_, err := db.Exec("INSERT INTO tasks (description, created_at, completed) VALUES (?, ?, ?)",
		task.Description, task.Created.Format(time.RFC3339), task.Completed,
	)
	return err
}

// GetAllTasksFromDB retrieves all tasks from the database
func GetAllTasksFromDB(db *sql.DB) ([]Task, error) {
	// Query the database for all tasks
	rows, err := db.Query("SELECT id, description, created_at, completed FROM tasks")
	if err != nil {
		log.Printf("error querying tasks: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Create a slice to hold the tasks
	var tasks []Task
	// Iterate over the rows and scan each task into a Task struct
	for rows.Next() {
		var task Task
		var created string
		if err := rows.Scan(&task.ID, &task.Description, &created, &task.Completed); err != nil {
			log.Printf("error scanning row: %v", err)
			return nil, err
		}
		task.Created, _ = time.Parse(time.RFC3339, created)
		tasks = append(tasks, task)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Printf("error scanning rows: %v", err)
		return nil, err
	}

	return tasks, nil
}

// UpdateTaskInDB updates a task in the database by ID
func UpdateTaskInDB(db *sql.DB, taskID int) error {
	// Update the task in the database
	_, err := db.Exec("UPDATE tasks SET completed = ? WHERE id = ?", true, taskID)
	if err != nil {
		fmt.Println("error updating task with ID %d: %v", taskID, err)
	}
	return err
}

// FormatTimeAsAgo formats a time.Time value as a human-readable "time ago" string
func FormatTimeAsAgo(created time.Time) string {
	duration := time.Since(created)
	if duration.Hours() >= 24 {
		days := int(duration.Hours() / 24)
		return fmt.Sprintf("%d days ago", days)
	} else if duration.Hours() >= 1 {
		hours := int(duration.Hours())
		return fmt.Sprintf("%d hours ago", hours)
	} else if duration.Minutes() >= 1 {
		minutes := int(duration.Minutes())
		return fmt.Sprintf("%d minutes ago", minutes)
	}
	return "just now"
}
