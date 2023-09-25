package schedule

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/shota-imoto/trvl/handler"
)

type ResponseBody struct {
	Schedule Schedule `json:"schedule"`
}

// テスト用のcurlコマンド
// curl -X POST -H "Content-Type: application/json" -d '{"start_date": "2011-01-07T17:44:13Z"}' http://localhost:5000/schedules
func CreateScheduleHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body struct {
		StartDate Date `json:"start_date" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handler.RespondJSON(ctx, w, &handler.ErrorResponse{
			Message: "CreateScheduleHandler",
			Details: []string{err.Error()},
		}, http.StatusInternalServerError)
		return
	}
	ss := NewService()

	s, err := ss.CreateSchedule(ctx, "userID", time.Time(body.StartDate))
	if err != nil {
		handler.RespondJSON(ctx, w, &handler.ErrorResponse{
			Message: "ss.CreateSchedule",
			Details: []string{err.Error()},
		}, http.StatusInternalServerError)
		return
	}
	handler.RespondJSON(ctx, w, ResponseBody{Schedule: s}, http.StatusOK)
}
