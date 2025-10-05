package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
	"github.com/programmingpercy/wedding/email"
	layouts "github.com/programmingpercy/wedding/templates"
	"github.com/programmingpercy/wedding/views/home"
	"github.com/programmingpercy/wedding/views/rsvp"
)

var emailService *email.EmailService

func main() {
	slog.Info("Welcome to Wedding of Pemiz")

	if err := godotenv.Load(); err != nil {
		slog.Info(fmt.Sprintf(".env not found, using system environment: %v", err))
	}

	es := email.NewEmailService()
	emailService = es

	fs := http.FileServer(http.Dir("./static/assets"))
	http.Handle("/static/assets/", http.StripPrefix("/static/assets/", fs))

	http.HandleFunc("/", Base)

	http.Handle("/osa", http.HandlerFunc(Osa))
	http.Handle("/speach", http.HandlerFunc(Speach))
	http.HandleFunc("/osa/people/new", OsaPeopleNew)
	http.HandleFunc("/osa/empty", HTMXEmpty)
	slog.Info("Starting on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Failed to host website", "error", err.Error())
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

type RSVPRequest struct {
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Coming  string   `json:"coming"`
	Message string   `json:"message"`
	People  []string `json:"people"`
}

func Osa(w http.ResponseWriter, r *http.Request) {
	slog.Info("New persons is attending")

	// Email the attendees to an configured email
	var req RSVPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Failed to decode RSVP JSON", "error", err.Error())

		if err := layouts.Popup("Något gick fel med anmälan, vänligen försök igen senare!", true).Render(r.Context(), w); err != nil {
			slog.Error("Failed to render success popup", "error", err.Error())
		}
		return
	}

	payload, err := convertRsvpRequestToEmail(req)

	if err != nil {
		slog.Error("Failed to convert rsvp to email req", "error", err)
		if err := layouts.Popup("Något gick fel med anmälan, vänligen försök igen senare!", true).Render(r.Context(), w); err != nil {
			slog.Error("Failed to render success popup", "error", err.Error())
		}
	}

	if err := emailService.Send([]string{"osa@pmwedding.se"}, "Ny osning", "", payload); err != nil {
		slog.Error(fmt.Sprintf("Failed to send email %w", err))
		// TODO: Print user info so we can still capture them
	}

	slog.Info("RSVP received", "name", req.Name, "email", req.Email, "count", len(req.People), "coming", req.Coming, "message", req.Message)

	if err := rsvp.OsaForm().Render(r.Context(), w); err != nil {
		slog.Error("Failed to render new OSA form", "error", err.Error())
	}
	if err := layouts.Popup("Tack för din anmälan!", false).Render(r.Context(), w); err != nil {
		slog.Error("Failed to render success popup", "error", err.Error())
	}
}

func convertRsvpRequestToEmail(rsvp RSVPRequest) (string, error) {
	var b strings.Builder
	if err := email.RsvpEmailTmpl.Execute(&b, rsvp); err != nil {
		return "", err
	}
	return b.String(), nil
}

type SpeachRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

func Speach(w http.ResponseWriter, r *http.Request) {
	slog.Info("New speach registered")

	// Email the attendees to an configured email
	var req SpeachRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Failed to decode Speach JSON", "error", err.Error())

		if err := layouts.Popup("Något gick fel med anmälan, vänligen försök igen senare!", true).Render(r.Context(), w); err != nil {
			slog.Error("Failed to render success popup", "error", err.Error())
		}
		return
	}

	slog.Info("Speach received", "name", req.Name, "email", req.Email, "phone", req.Phone, "message", req.Message)

	if err := rsvp.SpeachForm().Render(r.Context(), w); err != nil {
		slog.Error("Failed to render new OSA form", "error", err.Error())
	}
	if err := layouts.Popup("Tack för din anmälan, toastmasters hör av sig innan bröllopet för mer information!", false).Render(r.Context(), w); err != nil {
		slog.Error("Failed to render success popup", "error", err.Error())
	}
}

func OsaPeopleNew(w http.ResponseWriter, r *http.Request) {
	rsvp.PersonTile().Render(r.Context(), w)
}

func HTMXEmpty(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
