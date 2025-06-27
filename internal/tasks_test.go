package internal

import (
	"os"
	"testing"
)

func TestTaskCSVOperations(t *testing.T) {
	filename := "test_tasks.csv"
	defer os.Remove(filename)

	// Add tasks
	task1 := Task{ID: 1, Description: "Learn Go", Completed: false}
	task2 := Task{ID: 2, Description: "Code Go", Completed: false}
	task3 := Task{ID: 3, Description: "Master Go", Completed: false}

	if err := AppendTaskToCSV(task1, filename); err != nil {
		t.Fatalf("AppendTaskToCSV task1 failed: %v", err)
	}
	if err := AppendTaskToCSV(task2, filename); err != nil {
		t.Fatalf("AppendTaskToCSV task2 failed: %v", err)
	}
	if err := AppendTaskToCSV(task3, filename); err != nil {
		t.Fatalf("AppendTaskToCSV task3 failed: %v", err)
	}

	// Read tasks
	tasks, err := ReadTasksFromCSV(filename)
	if err != nil {
		t.Fatalf("ReadTasksFromCSV failed: %v", err)
	}
	if len(tasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(tasks))
	}

	// Update a task
	if err := UpdateTaskInCSV(task1.ID, filename); err != nil {
		t.Fatalf("UpdateTaskInCSV failed: %v", err)
	}
	tasks, err = ReadTasksFromCSV(filename)
	if err != nil {
		t.Fatalf("ReadTasksFromCSV after update failed: %v", err)
	}
	if !tasks[0].Completed {
		t.Errorf("Expected task1 to be completed, got %v", tasks[0].Completed)
	}
}
