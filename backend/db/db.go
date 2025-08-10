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
	// Use DSN to connect to the MySQL database.
	dsn := "root:my-secret-pw@tcp(my-mysql:3306)/goproject?parseTime=true"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// Test the connection to ensure it's working.
	if err = DB.Ping(); err != nil {
		DB.Close() // Close the connection if it's not valid.
		return err
	}

	// Set connection pool parameters.
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(time.Minute * 3)
	return nil
}

// CreateTable creates the users and events tables if they do not exist.
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
	if _, err := DB.Exec(createUsersTable); err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	createEventTable := `CREATE TABLE IF NOT EXISTS events (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		location VARCHAR(255) NOT NULL,
		startTime DATETIME NOT NULL,
		user_id INT,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`
	if _, err := DB.Exec(createEventTable); err != nil {
		log.Fatalf("Failed to create events table: %v", err)
	}
}
