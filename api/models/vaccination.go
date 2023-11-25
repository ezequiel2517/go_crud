package models

import (
	"time"
)

type Vaccination struct {
	ID     int
	Name   string
	DrugId int
	Dose   int
	Fecha  time.Time
}
