package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/shota-imoto/trvl/db"
	"github.com/shota-imoto/trvl/schedule"
)

func main() {
	fmt.Print("main")
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	// DB接続
	ctx := context.Background()
	err := db.Connect(ctx)
	if err != nil {
		return fmt.Errorf("run failed: %w", err)
	}

	// HTTPサーバー初期化
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	r.Route("/schedules", func(r chi.Router) {
		r.Post("/", schedule.CreateScheduleHandler)
	})
	err = http.ListenAndServe(":5000", r)
	if err != nil {
		return fmt.Errorf("run failed: %w", err)
	}
	return nil
}
