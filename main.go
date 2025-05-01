package main

import (
	"log/slog"
	"net/http"

	layouts "github.com/programmingpercy/wedding/templates"
)

func main() {
	slog.Info("Welcome to Wedding of Pemiz")

	fs := http.FileServer(http.Dir("./static/assets"))
	http.Handle("/static/assets/", http.StripPrefix("/static/assets/", fs))

	http.HandleFunc("/", Base)

	slog.Info("Starting on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Failed to host website", "error", err.Error())
	}
}

func Base(w http.ResponseWriter, r *http.Request) {
	slog.Info("New visitor", "method", r.Method)

	component := layouts.Base("test")
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render base", "error", err.Error())
	}
}

