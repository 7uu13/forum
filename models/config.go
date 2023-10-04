package models

import "database/sql"

func createTables(db *sql.DB) error {
	createUserTable := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT NOT NULL UNIQUE,
			age INTEGER NOT NULL DEFAULT 0,
			gender TEXT NOT NULL,
			firstname TEXT NOT NULL,
			lastname TEXT NOT NULL,
            email TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL
        );
    `

	createPostTable := `
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			user_id INTEGER,
			foreign key (user_id) REFERENCES users (id)
		);		
	`

	_, err := db.Exec(createUserTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(createPostTable)
	return err
}

func PerformMigrations(db *sql.DB) error {
	// Call your migration function here
	if err := createTables(db); err != nil {
		return err
	}

	// Other migrations can be added here if needed

	return nil
}
