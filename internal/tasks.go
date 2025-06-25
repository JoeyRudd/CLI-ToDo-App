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
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Create a record for each task
	record := []string{
		strconv.Itoa(task.ID),
		task.Description,
		strconv.FormatBool(task.Completed),
	}

	// Write the record to the CSV file
	if err := writer.Write(record); err != nil {
		return err
	}
	return nil
}

func ReadTasksFromCSV(filename string) ([]Task, error) {
	// create a slice to hold tasks
	var tasks []Task

	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		// Check if the record has the expected number of fields
		if len(record) != 3 {
			return nil, nil // or return an error if you prefer
		}
		// Convert the ID from string to int
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		// Convert the Completed status from string to bool
		completed, err := strconv.ParseBool(record[2])
		if err != nil {
			return nil, err
		}

		// Create a Task instance and append it to the slice
		task := Task{
			ID:          id,
			Description: record[1],
			Completed:   completed,
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
