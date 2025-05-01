package main

import (
	"log/slog"
	"net/http"

	layouts "github.com/programmingpercy/wedding/templates"
	"github.com/programmingpercy/wedding/views/faq"
	"github.com/programmingpercy/wedding/views/history"
	"github.com/programmingpercy/wedding/views/home"
	"github.com/programmingpercy/wedding/views/rsvp"
	"github.com/programmingpercy/wedding/views/weddingday"
)

func main() {
	slog.Info("Welcome to Wedding of Pemiz")

	fs := http.FileServer(http.Dir("./static/assets"))
	http.Handle("/static/assets/", http.StripPrefix("/static/assets/", fs))

	http.HandleFunc("/", Base)
	http.Handle("/home", LayoutMiddleware(http.HandlerFunc(Home)))
	http.Handle("/weddingday", LayoutMiddleware(http.HandlerFunc(WeddingDay)))
	http.HandleFunc("/history", History)
	http.HandleFunc("/rsvp", Rsvp)
	http.HandleFunc("/faq", Faq)
	http.HandleFunc("/crew", Crew)

	slog.Info("Starting on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Failed to host website", "error", err.Error())
	}
}

func LayoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") == "true" {
			next.ServeHTTP(w, r)
		} else {
			// Full page reload, so print layout also
			home := home.Home()
			layout := layouts.Base("Test", home)

			if err := layout.Render(r.Context(), w); err != nil {
				slog.Error("Failed to render layout", "error", err.Error())
			}
		}
	})
}

func Base(w http.ResponseWriter, r *http.Request) {
	slog.Info("New visitor", "method", r.Method)

	base := home.Home()
	component := layouts.Base("test", base)
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render base", "error", err.Error())
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	slog.Info("New visitor", "page", "home")

	component := home.Home()
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render home component", "error", err.Error())
	}
}

func WeddingDay(w http.ResponseWriter, r *http.Request) {
	slog.Info("New visitor", "page", "weddingday")

	component := weddingday.WeddingDay()
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render weddingday component", "error", err.Error())
	}
}

func History(w http.ResponseWriter, r *http.Request) {
	slog.Info("New visitor", "page", "history")

	component := history.History()
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render history component", "error", err.Error())
	}
}

func Rsvp(w http.ResponseWriter, r *http.Request) {
	slog.Info("New visitor", "page", "rsvp")

	component := rsvp.RSVP()
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render rsvp component", "error", err.Error())
	}
}

func Faq(w http.ResponseWriter, r *http.Request) {
	slog.Info("New visitor", "page", "faq")

	component := faq.Faq()
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render faq component", "error", err.Error())
	}
}

func Crew(w http.ResponseWriter, r *http.Request) {
	slog.Info("New visitor", "page", "crew")

	component := faq.Faq()
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render crew component", "error", err.Error())
	}
}
