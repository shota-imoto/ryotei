package schedule

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shota-imoto/trvl/db"
)

type ScheduleService struct {
	Query *db.Queries
}

func NewService() ScheduleService {
	return ScheduleService{Query: db.Connection}
}

func (ss ScheduleService) CreateSchedule(ctx context.Context, userID string, time time.Time) (Schedule, error) {
	schedule, err := ss.Query.CreateSchedule(ctx, db.CreateScheduleParams{
		UserID:    userID,
		StartDate: pgtype.Timestamp{Time: time, Valid: true},
	})
	if err != nil {
		return Schedule{}, fmt.Errorf("ss.Query.CreateSchedule is failed: %w", err)
	}
	s := FromDB(schedule)
	return s, nil
}
