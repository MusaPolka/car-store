package email

import (
	"context"
	"fmt"
	_ "github.com/emersion/go-sasl"
	_ "go/token"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	_ "log"
	"net/smtp"
	"os"
)

func SendEmail(to, subject, body string) error {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/gmail.send"},
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/callback",
	}

	// Получение и обновление токена
	token := &oauth2.Token{RefreshToken: os.Getenv("GOOGLE_REFRESH_TOKEN")}
	ts := conf.TokenSource(context.Background(), token)
	currentToken, err := ts.Token()
	if err != nil {
		return fmt.Errorf("токен недействителен: %v", err)
	}

	// Исправленная аутентификация
	auth := &oauth2.Transport{
		Source: conf.TokenSource(context.Background(), currentToken),
	}

	msg := []byte(fmt.Sprintf(
		"To: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
		to, subject, body,
	))

	return smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("SMTP_USER"),
		[]string{to},
		msg,
	)
}
