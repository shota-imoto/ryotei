package schedule_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/shota-imoto/trvl/db"
	"github.com/shota-imoto/trvl/schedule"
)

func TestCreateSchedule(t *testing.T) {
	ctx := context.Background()
	db.Connect(ctx)

	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer tx.Rollback(ctx)
	txQuery := db.Connection.WithTx(tx)

	now := time.Now()
	_, err = schedule.ScheduleService{Query: txQuery}.CreateSchedule(ctx, "userID", now.UTC())
	if err != nil {
		t.Errorf("CreateSchedule is failed: %v", err)
	}

	schedules, err := txQuery.ListSchedule(ctx)
	if err != nil {
		t.Errorf("db.Connection.ListSchedule is failed: %v", err)
	}

	if len(schedules) != 1 {
		t.Errorf("len(schedules) expect: %d, got: %d", 1, len(schedules))
	}

	if schedules[0].UserID != "userID" {
		t.Errorf("schedules[0].UserID expect: %s, got: %s", "userID", schedules[0].UserID)
	}

	got := schedules[0].StartDate.Time.String()
	if diff := cmp.Diff(got, now.UTC().String()); diff != "" {
		fmt.Printf("%v", got)
		t.Errorf("-schedules[0].StartDate.Time, +now: %s\n", diff)
	}

}
