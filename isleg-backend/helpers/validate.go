package helpers

import (
	"regexp"
)

// Verify e-mail address
func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func ValidatePhoneNumber(regPersonalNumber string) bool {
	regexpPersonalNumber := regexp.MustCompile(`^(\+9936)[1-5][0-9]{6}$`)
	isMatchPersonalNumber := regexpPersonalNumber.MatchString(regPersonalNumber)
	return isMatchPersonalNumber
}
