package models

import (
	"example.com/rest-api/db"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" bind:"required,email"`
	Username string `json:"username" bind:"required"`
	Password string `json:"password" bind:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, username, password) VALUES (?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Username, u.Password)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = id

	return nil
}	