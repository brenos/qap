package helpers

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func ValidateEmail(email string) error {
	emailSplited := strings.Split(email, "@")
	domainError := validateEmailDomain(emailSplited)
	if domainError != nil {
		return domainError
	}
	domain := strings.ToLower(emailSplited[len(emailSplited)-1])
	return isDisposableEmail(domain)
}

func isDisposableEmail(domain string) error {
	f, _ := os.Open("blacklistEmails.conf")
	for scanner := bufio.NewScanner(f); scanner.Scan(); {
		if scanner.Text() == domain {
			f.Close()
			return errors.New("Email is not valid!")
		}
	}
	f.Close()
	return nil
}

func validateEmailDomain(domains []string) error {
	if len(domains) != 2 {
		return errors.New("Error on email validation!")
	}
	return nil
}
