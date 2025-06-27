package internal

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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

// AppendTaskToCSV appends a new task to the CSV file
func AppendTaskToCSV(task Task, filename string) error {
	// Open the file for writing
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer CloseFile(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Create a record for each task
	record := []string{
		strconv.Itoa(task.ID),
		task.Description,
		task.Created.Format(time.RFC3339),
		strconv.FormatBool(task.Completed),
	}

	// Write the record to the CSV file
	if err := writer.Write(record); err != nil {
		return err
	}
	return nil
}

// ReadTasksFromCSV reads tasks from a CSV file and returns a slice of Task
func ReadTasksFromCSV(filename string) ([]Task, error) {
	// create a slice to hold tasks
	var tasks []Task

	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer CloseFile(file)

	// Create a CSV reader
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		// Check if the record has the expected number of fields
		if len(record) != 4 {
			return nil, fmt.Errorf("malformed record at line %d: %v", len(tasks)+1, record)
		}
		// Convert the ID from string to int
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("invalid ID in record %v: %v", record, err)
		}
		created, err := time.Parse(time.RFC3339, record[2])
		if err != nil {
			return nil, fmt.Errorf("invalid time in record %v: %v", record, err)
		}
		// Convert the Completed status from string to bool
		completed, err := strconv.ParseBool(record[3])
		if err != nil {
			return nil, err
		}

		// Create a Task instance and append it to the slice
		task := Task{
			ID:          id,
			Description: record[1],
			Created:     created,
			Completed:   completed,
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// UpdateTaskInCSV updates a task in the CSV file by ID
func UpdateTaskInCSV(taskID int, filename string) error {
	// Read all tasks from the CSV file
	tasks, err := ReadTasksFromCSV(filename)
	if err != nil {
		return err
	}

	// Find the task to update
	found := false
	for i, t := range tasks {
		if t.ID == taskID {
			tasks[i].Completed = true // Update the task
			found = true
			break
		}
	}

	// If the task was not found, return an error
	if !found {
		return fmt.Errorf("task with ID %d not found", taskID)
	}

	// Open the file for writing (overwrite)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer CloseFile(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write updated tasks back to the CSV file
	for _, t := range tasks {
		record := []string{
			strconv.Itoa(t.ID),
			t.Description,
			t.Created.Format(time.RFC3339),
			strconv.FormatBool(t.Completed),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

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
