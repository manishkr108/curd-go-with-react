package models

import (
	"backend/db"
	"backend/utils"
	"errors"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u User) Save() error {
	query := "INSERT INTO users(name,email,password)  VALUE(?,?,?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	hashpass, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Name, u.Email, hashpass)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {

	query := "SELECT id, password FROM users WHERE email=?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string

	err := row.Scan(&u.ID, &retrivedPassword)

	if err != nil {
		return errors.New("credintials Invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivedPassword)

	if !passwordIsValid {
		return errors.New("invalid password")
	}

	return nil
}
