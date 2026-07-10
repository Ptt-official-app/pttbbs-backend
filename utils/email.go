package utils

import (
	"encoding/base64"
	"fmt"
	"mime"
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/Ptt-official-app/pttbbs-backend/types"
)

// https://gist.github.com/andelf/5004821
func SendEmail(rcpts []string, title string, content string) (err error) {
	if isTest {
		return nil
	}

	from := mail.Address{Name: types.BBSNAME + types.SENDER_SUFFIX, Address: types.EMAIL_FROM}
	to := make([]string, len(rcpts))
	for idx, each := range rcpts {
		eachAddr := mail.Address{Name: "", Address: each}
		to[idx] = eachAddr.String()
	}

	subject := encodeRFC2047(title)
	header := map[string]string{
		"From":                      from.String(),
		"To":                        strings.Join(to, ", "),
		"Subject":                   subject,
		"MIME-Version":              "1.0",
		"Content-Type":              "text/html; charset=\"utf-8\"",
		"Content-Transfer-Encoding": "base64",
	}

	message := ""

	for k, v := range header {
		message += fmt.Sprintf("%s: %s\n", k, v)
	}
	message += "\n" + base64.StdEncoding.EncodeToString([]byte(content))

	err = smtp.SendMail(types.EMAIL_SERVER, nil, from.Address, rcpts, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

func encodeRFC2047(str string) string {
	// use mail's rfc2047 to encode any string
	return mime.BEncoding.Encode("utf-8", str)
}
