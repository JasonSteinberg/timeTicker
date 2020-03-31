package users

import (
	"database/sql"
	"errors"
	"github.com/JasonSteinberg/timeTicker/database"
	"github.com/JasonSteinberg/timeTicker/structs"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreateUser(user structs.User) error {
	db := database.GetSqlWriteDB()

	result, err := db.Exec(`Insert into user (email, password) values (?,?);`, user.Email, user.Password)
	if err != nil {
		return errors.New("Unable to create user! " + err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return errors.New("Unable to get id of user! " + err.Error())
	}
	user.ID = id

	return nil
}

func CheckUser(userAttempted structs.User, password string) (bool, error) {
	var user structs.User
	db := database.GetSqlReadDB()

	row, err := db.Query("select id, email, password, access from user where email=?", userAttempted.Email)
	row.Next()
	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("The user does not exist")
		} else {
			log.Fatal(err)
			return false, err
		}
	}

	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Access); err != nil {
		log.Fatal(err)
		return false, err

	}

	hashedPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
