package utils

import "net/mail"

func ValidEmail(email string) (valid bool) {
	_, err := mail.ParseAddress(email)
	return err == nil
}
