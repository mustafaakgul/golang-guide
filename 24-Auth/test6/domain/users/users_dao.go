package users

import (
	"fmt"
	"test6/datasource"
	"test6/utils/errors"
)

var (
	queryInsertUser     = "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?);"
	queryGetUserByEmail = "SELECT id, first_name, last_name, email, password FROM users WHERE email=?;"
	queryGetUserByID    = "SELECT id, first_name, last_name, email FROM users WHERE id=?;"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := datasource.dbx.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if saveErr != nil {
		fmt.Println(saveErr)
		return errors.NewInternalServerError("database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	user.ID = userID
	return nil
}

func (user *User) GetByEmail() *errors.RestErr {
	stmt, err := datasource.Client.Prepare(queryGetUserByEmail)
	if err != nil {
		return errors.NewInternalServerError("invalid email")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password); getErr != nil {
		fmt.Println(getErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) GetByID() *errors.RestErr {
	stmt, err := datasource.Client.Prepare(queryGetUserByID)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); getErr != nil {
		return errors.NewInternalServerError("database error")
	}
	return nil
}
