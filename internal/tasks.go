package internal

import (
	"encoding/csv"
	"os"
	"strconv"
)

// create a struct for a task
type Task struct {
	ID          int
	Description string
	Completed   bool
}

// create a slice to hold tasks
func (t Task) addTask(description string) Task {
	return Task{
		ID:          t.ID + 1, // Increment ID for simplicity
		Description: description,
		Completed:   false,
	}
}

func AppendTaskToCSV(task Task, filename string) error {
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// create a record for each task
	record := []string{
		strconv.Itoa(task.ID),
		task.Description,
		strconv.FormatBool(task.Completed),
	}

	// write the record to the CSV file
	if err := writer.Write(record); err != nil {
		return err
	}

	return nil

}
