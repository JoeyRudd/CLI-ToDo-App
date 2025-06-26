package internal

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

// CloseFile Helper function to close a file and log an error if it occurs
func CloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Printf("error closing file: %v", err)
	}
}

// Create a slice to hold tasks
func (t Task) addTask(description string) Task {
	return Task{
		ID:          t.ID + 1, // Increment ID for simplicity
		Description: description,
		Completed:   false,
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
		if len(record) != 3 {
			return nil, fmt.Errorf("Malformed record at line %d: %v", len(tasks)+1, record)
		}
		// Convert the ID from string to int
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("Invalid ID in record %v: %v", record, err)
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

// UpdateTaskInCSV updates a task in the CSV file by ID
func UpdateTaskInCSV(task Task, filename string) error {
	// Read all tasks from the CSV file
	tasks, err := ReadTasksFromCSV(filename)
	if err != nil {
		return err
	}

	// Find the task to update
	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i].Completed = true // Update the task
			break
		}
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
			strconv.FormatBool(t.Completed),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
