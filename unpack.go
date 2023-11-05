package unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

type state int

const (
	start state = iota
	escape
	number
)

func Unpack(str string) (string, error) {
	var currentState state = start
	var result strings.Builder
	numOfRepeat := 1
	//if firstCheckString(&str) == true {
	//	return "", ErrInvalidString
	//}

	var runeArray = []rune(str)

	if len(runeArray) > 0 && unicode.IsDigit(runeArray[0]) {
		return "first rune is digit", ErrInvalidString
	}

	for i, char := range runeArray {

		switch currentState {
		case start:
			if unicode.IsDigit(char) {
				currentState = number
				numOfRepeat, _ = strconv.Atoi(string(char))
			} else if char == '\\' {
				currentState = escape
			} else {
				currentState = start
				result.WriteRune(char)
			}
		case number:
			if unicode.IsDigit(char) {
				return "", ErrInvalidString
			}
			repStr, err := repeatRune(runeArray[i-2], numOfRepeat)
			if err != nil {
				res := result.String()
				result.Reset()
				result.WriteString(res[0 : len(res)-1])
			}
			result.WriteString(repStr)
			currentState = start
			result.WriteRune(char)
		case escape:
			result.WriteRune(char)

		}

	}

	return result.String(), nil

}

func repeatRune(char rune, numOfRepeat int) (string, error) {
	if numOfRepeat > 0 {
		return strings.Repeat(string(char), numOfRepeat-1), nil
	}
	if numOfRepeat == 0 {
		return "", errors.New("zero repeat")
	}
	return "", nil
}
