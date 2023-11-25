package usercontroller

import (
	"api/api/models"
	"api/db"
)

func InsertUser(user models.User) error {
	db := db.GetConnection()
	sqlStatement := `
	INSERT INTO users (name, email, password)
	VALUES ($1, $2, $3)
	`

	_, err := db.Exec(sqlStatement, user.Name, user.Email, user.Password)
	return err
}

func GetUser(email string) (models.User, error) {
	db := db.GetConnection()
	sqlStatement := `
	SELECT name, email, password
	FROM users
	WHERE email = $1
	`

	row := db.QueryRow(sqlStatement, email)
	var user models.User
	err := row.Scan(&user.Name, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
