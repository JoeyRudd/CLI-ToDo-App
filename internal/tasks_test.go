package internal

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
	"os"
	"testing"
)

// SetupTestDB initializes an in-memory SQLite database for testing purposes.
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	_, err = db.Exec(`CREATE TABLE tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT,
		created_at TEXT,
		completed BOOLEAN
	)`)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}
	return db
}

func TestTaskSQLOperations(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Add tasks
	task1 := Task{Description: "Learn SQL", Completed: false}
	task2 := Task{Description: "Code SQL", Completed: false}
	task3 := Task{Description: "Master SQL", Completed: false}

	if err := AddTaskToDB(db, task1); err != nil {
		t.Fatalf("AddTaskToDB task1 failed: %v", err)
	}
	if err := AddTaskToDB(db, task2); err != nil {
		t.Fatalf("AddTaskToDB task2 failed: %v", err)
	}
	if err := AddTaskToDB(db, task3); err != nil {
		t.Fatalf("AddTaskToDB task3 failed: %v", err)
	}

	// Read tasks
	tasks, err := GetAllTasksFromDB(db)
	if err != nil {
		t.Fatalf("GetTasksFromDB failed: %v", err)
	}
	if len(tasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(tasks))
	}

	// Update a task
	if err := UpdateTaskInDB(db, 1); err != nil {
		t.Fatalf("UpdateTaskInDB failed: %v", err)
	}
	tasks, err = GetAllTasksFromDB(db)
	if err != nil {
		t.Fatalf("GetTasksFromDB after update failed: %v", err)
	}
	if !tasks[0].Completed {
		t.Errorf("Expected task1 to be completed, got %v", tasks[0].Completed)
	}
}
