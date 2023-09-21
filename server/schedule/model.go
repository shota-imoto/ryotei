package schedule

import (
	"time"

	"github.com/shota-imoto/trvl/db"
)

type Schedule struct {
	UserID    string    `json:"user_id"`
	StartDate time.Time `json:"start_date"`
}

func FromDB(db db.Schedule) Schedule {
	s := Schedule{
		UserID:    db.UserID,
		StartDate: db.StartDate.Time,
	}
	return s
}
