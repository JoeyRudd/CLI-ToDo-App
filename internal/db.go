package internal

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

var DB *sql.DB // Exported DB variable

func InitDB(dbPath string) (*sql.DB, error) {
	// Open the database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	sql := `
	CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed BOOLEAN DEFAULT FALSE
	);
	`

	_, err = db.Exec(sql)
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
