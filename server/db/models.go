// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Schedule struct {
	ID        int32
	UserID    string
	StartDate pgtype.Timestamp
	CreatedAt pgtype.Timestamp
}
