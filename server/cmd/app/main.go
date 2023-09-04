package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/shota-imoto/trvl"
)

func main() {
	fmt.Print("main")
	run()
}

func run() {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	trvl.Handle()

	http.ListenAndServe(":5000", r)
}
