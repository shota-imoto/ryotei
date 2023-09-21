package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Connection *Queries
var Pool *pgxpool.Pool

// DB接続インスタンスを生成し、Connectionに格納する
func Connect(ctx context.Context) error {
	var err error
	uri, err := uri()
	if err != nil {
		return fmt.Errorf("db.connect failed: %w", err)
	}
	Pool, err = pgxpool.New(ctx, uri)
	if err != nil {
		return fmt.Errorf("db connect failed: %w", err)
	}
	Connection = New(Pool)
	return nil
}

func uri() (string, error) {
	env := os.Getenv("ENV")
	if env == "" {
		return "", fmt.Errorf("環境名が指定されていません。dev, prod, stgのいずれかを指定してください")
	}
	if env == "dev" {
		return "postgres://postgres:postgres@db:5432/trvl?sslmode=disable", nil
	}
	if env == "test" {
		return "postgres://postgres:postgres@localhost:5432/trvl_test?sslmode=disable", nil
	}
	return "", fmt.Errorf("不正な環境名")
}
