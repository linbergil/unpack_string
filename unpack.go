package unpackstring

import (
	"errors"
	"github.com/GRbit/go-pcre"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	// Place your code here.

	re, _ := pcre.Compile(`^\d|\d{2,}|\\(?=\D)|(?=\\)`, pcre.UTF8)

	res := re.MatchStringWFlags(str, pcre.UTF8)

	if res == true {
		return "", ErrInvalidString
	}

	return "", nil
}
