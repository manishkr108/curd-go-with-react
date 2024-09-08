package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitDB initializes the database connection.
func InitDB() error {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/goproject?parseTime=true")
	if err != nil {
		return err
	}

	// Test the connection to ensure it's working
	if err = DB.Ping(); err != nil {
		DB.Close() // Close the connection if it's not valid
		return err
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(time.Minute * 3)
	return nil
}

// CreateTable creates the events table if it does not exist.
func CreateTable() {
	if DB == nil {
		log.Fatal("Database connection is not initialized")
	}

	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password TEXT NOT NULL
		
	)`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("could not create users table")
	}

	createEventTable := `CREATE TABLE IF NOT EXISTS events (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		location VARCHAR(255) NOT NULL,
		startTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventTable)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}
