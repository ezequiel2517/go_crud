package models

import (
	"database/sql"
	"time"
)

type Drug struct {
	ID          int
	Name        string
	Approved    bool
	MinDose     int
	MaxDose     int
	AvailableAt time.Time
}

func InsertarDrug(db *sql.DB, name string, approved bool,
	min_dose int, max_dose int, available_at time.Time) error {
	sqlStatement := `
    INSERT INTO drugs (name, approved, min_dose, max_dose, available_at)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id`

	var drugId int

	err := db.QueryRow(sqlStatement, name, approved, min_dose, max_dose, available_at).Scan(&drugId)
	if err != nil {
		return err
	}

	return nil
}
