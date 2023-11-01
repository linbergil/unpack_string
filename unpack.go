package unpackstring

import (
	"errors"
	"regexp"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	// Place your code here.
	err := CheckString(str)
	if err != nil {
		return "", err
	}

	return "", nil
}

func CheckString(str string) error {

	re, _ := regexp.Compile(`^\d|\d{2,}`)

	if re.MatchString(str) == true {
		return ErrInvalidString
	}

	return nil
}
