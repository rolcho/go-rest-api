package models

import (
	"database/sql"
	"time"

	"github.com/rolcho/go-rest-api/db"
)

type User struct {
	Id        int64
	Name      string
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Save() error {
	query := `
	INSERT INTO users (name, email, password, createdAt, updatedAt)
	VALUES (?, ?, ?, ?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	u.CreatedAt = time.Now().Truncate(time.Second)
	sqlCreatedAt := u.CreatedAt.Format("2006-01-02T15:04:05Z")
	u.UpdatedAt = time.Now().Truncate(time.Second)
	sqlUpdatedAt := u.UpdatedAt.Format("2006-01-02T15:04:05Z")

	result, err := statement.Exec(u.Name, u.Email, u.Password, sqlCreatedAt, sqlUpdatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.Id = id

	return err
}

func GetAllUsers() ([]User, error) {
	query := `
	SELECT * FROM users
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email,
			&user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func GetUserById(id int64) (*User, error) {
	query := `
	SELECT * FROM users WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)

	var user User
	if err := row.Scan(&user.Id, &user.Name, &user.Email,
		&user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	query := `
	SELECT * FROM users WHERE email = ?
	`
	row := db.DB.QueryRow(query, email)

	var user User
	if err := row.Scan(&user.Id, &user.Name, &user.Email,
		&user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *User) Update() error {
	query := `
	UPDATE users
	SET name = ?, email = ?, password = ?, updatedAt = ?
	WHERE id = ?
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	u.UpdatedAt = time.Now().Truncate(time.Second)
	sqlUpdatedAt := u.UpdatedAt.Format("2006-01-02T15:04:05Z")

	_, err = statement.Exec(u.Name, u.Email, u.Password, sqlUpdatedAt, u.Id)

	return err
}

func (u *User) Delete() error {
	query := `
	DELETE FROM users
	WHERE id = ?
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(u.Id)

	return err
}
