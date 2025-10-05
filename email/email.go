package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"mime"
	"mime/multipart"
	"mime/quotedprintable"
	"net"
	"net/mail"
	"net/smtp"
	"net/textproto"
	"os"
	"strconv"
	"time"
)

type EmailService struct {
	Host      string
	Port      int
	Username  string
	Password  string
	FromName  string
	FromEmail string
	Timeout   time.Duration
}

// NewEmailService builds a service from environment variables.
// Defaults to one.com’s SMTP settings (implicit TLS on 465).
func NewEmailService() *EmailService {
	port := 465
	if p := os.Getenv("SMTP_PORT"); p != "" {
		if n, err := strconv.Atoi(p); err == nil {
			port = n
		}
	}
	return &EmailService{
		Host:      os.Getenv("SMTP_HOST"),
		Port:      port,
		Username:  os.Getenv("SMTP_USERNAME"),
		Password:  os.Getenv("SMTP_PASSWORD"),
		FromName:  os.Getenv("SMTP_FROM_NAME"),
		FromEmail: os.Getenv("SMTP_FROM_EMAIL"),
		Timeout:   10 * time.Second,
	}
}

// Send sends an email using SMTP with implicit TLS (port 465).
func (s EmailService) Send(to []string, subject, textBody, htmlBody string) error {
	if s.Host == "" || s.Username == "" || s.Password == "" || s.FromEmail == "" {
		return fmt.Errorf("email service not configured (missing SMTP env vars)")
	}

	from := mail.Address{Name: s.FromName, Address: s.FromEmail}
	allRcpts := append([]string{}, to...)

	// Build headers
	var msg bytes.Buffer
	msg.WriteString(fmt.Sprintf("Date: %s\r\n", time.Now().Format(time.RFC1123Z)))
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString(fmt.Sprintf("From: %s\r\n", from.String()))
	if len(to) > 0 {
		msg.WriteString(fmt.Sprintf("To: %s\r\n", to))
	}
	// UTF-8 subject
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", mime.QEncoding.Encode("utf-8", subject)))

	// Start multipart/alternative
	alt := multipart.NewWriter(&msg)
	msg.WriteString(fmt.Sprintf("Content-Type: multipart/alternative; boundary=%q\r\n", alt.Boundary()))
	msg.WriteString("\r\n") // end of headers

	// text/plain part
	if textBody != "" {
		h := textproto.MIMEHeader{}
		h.Set("Content-Type", "text/plain; charset=utf-8")
		h.Set("Content-Transfer-Encoding", "quoted-printable")
		w, _ := alt.CreatePart(h)
		qp := quotedprintable.NewWriter(w)
		_, _ = qp.Write([]byte(textBody))
		_ = qp.Close()
	}

	// text/html part
	if htmlBody != "" {
		h := textproto.MIMEHeader{}
		h.Set("Content-Type", "text/html; charset=utf-8")
		h.Set("Content-Transfer-Encoding", "quoted-printable")
		w, _ := alt.CreatePart(h)
		qp := quotedprintable.NewWriter(w)
		_, _ = qp.Write([]byte(htmlBody))
		_ = qp.Close()
	}

	_ = alt.Close()

	// Connect with implicit TLS (465)
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	dialer := &net.Dialer{Timeout: s.Timeout}
	tlsCfg := &tls.Config{ServerName: s.Host}

	conn, err := tls.DialWithDialer(dialer, "tcp", addr, tlsCfg)
	if err != nil {
		return fmt.Errorf("tls dial: %w", err)
	}
	defer conn.Close()

	c, err := smtp.NewClient(conn, s.Host)
	if err != nil {
		return fmt.Errorf("new smtp client: %w", err)
	}
	defer c.Quit()

	// Auth
	if ok, _ := c.Extension("AUTH"); ok {
		auth := smtp.PlainAuth("", s.Username, "F^'Z97$LtK5qyJr", s.Host)
		if err := c.Auth(auth); err != nil {
			return fmt.Errorf("smtp auth: %w", err)
		}
	}

	// MAIL FROM / RCPT TO
	if err := c.Mail(from.Address); err != nil {
		return fmt.Errorf("MAIL FROM: %w", err)
	}
	for _, rcpt := range allRcpts {
		if rcpt == "" {
			continue
		}
		if err := c.Rcpt(rcpt); err != nil {
			return fmt.Errorf("RCPT TO %s: %w", rcpt, err)
		}
	}

	// DATA
	w, err := c.Data()
	if err != nil {
		return fmt.Errorf("DATA: %w", err)
	}
	if _, err := w.Write(msg.Bytes()); err != nil {
		return fmt.Errorf("write body: %w", err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("close data: %w", err)
	}

	return nil
}
