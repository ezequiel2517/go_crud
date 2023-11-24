package models

import "database/sql"

func InsertarUsuario(db *sql.DB, name, email, password string) error {
	sqlStatement := `
    INSERT INTO usuario (name, email, password)
    VALUES ($1, $2, $3)
    RETURNING id`

	var userID int

	err := db.QueryRow(sqlStatement, name, email, password).Scan(&userID)
	if err != nil {
		return err
	}

	return nil
}
