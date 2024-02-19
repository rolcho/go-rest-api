package models

import (
	"time"

	"github.com/rolcho/go-rest-api/db"
)

type Event struct {
	Id          int64
	Title       string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserId      int
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id, createdAt, updatedAt)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	e.CreatedAt = time.Now().Truncate(time.Second)
	sqlCreatedAt := e.CreatedAt.Format("2006-01-02T15:04:05Z")
	e.UpdatedAt = time.Now().Truncate(time.Second)
	sqlUpdatedAt := e.UpdatedAt.Format("2006-01-02T15:04:05Z")

	result, err := statement.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserId, sqlCreatedAt, sqlUpdatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.Id = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT * FROM events
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.Id, &event.Title, &event.Description,
			&event.Location, &event.DateTime, &event.CreatedAt, &event.UpdatedAt, &event.UserId); err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `
	SELECT * FROM events WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)

	var event Event
	if err := row.Scan(&event.Id, &event.Title, &event.Description,
		&event.Location, &event.DateTime, &event.CreatedAt, &event.UpdatedAt, &event.UserId); err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?, updatedAt = ?
	WHERE id = ?
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	e.UpdatedAt = time.Now().Truncate(time.Second)
	sqlUpdatedAt := e.UpdatedAt.Format("2006-01-02T15:04:05Z")

	_, err = statement.Exec(e.Title, e.Description, e.Location, e.DateTime, sqlUpdatedAt, e.Id)

	return err
}

func (e *Event) Delete() error {
	query := `
	DELETE FROM events
	WHERE id = ?
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(e.Id)

	return err
}
