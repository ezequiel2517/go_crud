package models

import (
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
