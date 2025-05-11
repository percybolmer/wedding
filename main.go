package main

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	layouts "github.com/programmingpercy/wedding/templates"
	"github.com/programmingpercy/wedding/views/crew"
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
	http.Handle("/history", LayoutMiddleware(http.HandlerFunc(History)))
	http.Handle("/rsvp", LayoutMiddleware(http.HandlerFunc(Rsvp)))
	http.Handle("/faq", LayoutMiddleware(http.HandlerFunc(Faq)))
	http.Handle("/crew", LayoutMiddleware(http.HandlerFunc(Crew)))

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
			//
			path := r.URL.Path

			comp := getComponentFromPath(path)
			layout := layouts.Base("Test", comp)

			slog.Info("Path", "path", path)
			if err := layout.Render(r.Context(), w); err != nil {
				slog.Error("Failed to render layout", "error", err.Error())
			}
		}
	})
}

func getComponentFromPath(path string) templ.Component {
	switch path {
	case "/history":
		return history.History()
	case "/weddingday":
		return weddingday.WeddingDay()
	case "/rsvp":
		return rsvp.RSVP()
	case "/faq":
		return faq.Faq()
	case "/crew":
		return crew.Crew()
	default:
		return home.Home()
	}
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

	component := crew.Crew()
	if err := component.Render(r.Context(), w); err != nil {
		slog.Error("Failed to render crew component", "error", err.Error())
	}
}
