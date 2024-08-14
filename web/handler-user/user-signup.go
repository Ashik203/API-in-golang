package handleruser

import (
	"app/db"
	"app/logger"
	"app/web/utils"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/smtp"
	"time"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	log.Printf("Requesting to create user.")
	var user db.User

	json.NewDecoder(r.Body).Decode(&user)

	err := db.SignUp(&user)
	if err != nil {
		slog.Error("error in singing up", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": user,
		}))
		utils.SendError(w, http.StatusBadRequest, err.Error(), user)
		return
	}

	err = SendEmail(&user)
	if err != nil {
		slog.Error("error in sending email", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": user,
		}))
		utils.SendError(w, http.StatusInternalServerError, err.Error(), user)
		return
	}

	utils.SendData(w, user)
}

func GenerateSecureOTP(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}

	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func SendEmail(b *db.User) error {
	from := "mdfazlerabbispondontechnonext@gmail.com"
	password := "anyuuzkchpawhzcl"

	to := []string{
		b.Email,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	otp := GenerateSecureOTP(6)

	t, err := template.ParseFiles("/home/ashikurrahman/Documents/API-in-golang/html/signup-email.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Confirm SignUP \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Otp string
	}{
		Otp: otp,
	})

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return err
	}

	err = db.GetRedisClient().Set(context.Background(), b.Email, otp, 5*time.Minute).Err()
	if err != nil {
		return err
	}

	fmt.Println("Email Sent Successfully!")
	return err
}
