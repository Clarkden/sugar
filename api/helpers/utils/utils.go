package utils

import (
	"net/mail"
	"regexp"
	"strings"
)

func ValidEmail(email string) (valid bool) {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidDomain(domain string) bool {
	const domainRegex = `^(?i)(?:[a-z0-9](?:[a-z0-9\-]{0,61}[a-z0-9])?\.)+[a-z]{2,}$`

	re := regexp.MustCompile(domainRegex)

	if len(domain) < 1 || len(domain) > 253 {
		return false
	}

	domain = strings.TrimSuffix(domain, ".")

	return re.MatchString(domain)
}
