-- name: CreateSchedule :one
INSERT INTO schedule (
  user_id, start_date
) VALUES (
  $1, $2
)
RETURNING *;

-- name: ListSchedule :many
SELECT * FROM schedule ORDER BY id;