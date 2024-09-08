package models

import (
	"backend/db"
	"log"
	"time"
)

type Event struct {
	ID          int64     // id of the event
	Name        string    `json:"name" binding:"required"` // name of the event
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`  // location of the event
	StartTime   time.Time `json:"startTime" binding:"required"` // start time of the event
	UserID      int       `json:"user_id"`
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, startTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.StartTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT id, name, description, location, startTime, user_id FROM events"
	rows, err := db.DB.Query(query)
	log.Print(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event

		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.StartTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.StartTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

// Update updates an existing event in the database.
func (event Event) Update() error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, startTime = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.StartTime, event.ID)
	return err
}
