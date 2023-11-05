package unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// ErrInvalidString - Определение пользовательской ошибки для недопустимой строки.
var ErrInvalidString = errors.New("invalid string")

// state - Тип-перечисление для представления состояния автомата.
type state int

const (
	start state = iota
	escape
	number
)

// Unpack - Функция для распаковки строки в заданном формате.
func Unpack(str string) (string, error) {
	var currentState state = start
	var result strings.Builder
	numOfRepeat := 1

	var runeArray = []rune(str)

	if len(runeArray) > 0 && unicode.IsDigit(runeArray[0]) {
		// Если первый символ является цифрой, возвращаем ошибку "недопустимой строки".
		return "first rune is digit", ErrInvalidString
	}

	for i, char := range runeArray {

		switch currentState {
		case start:
			if unicode.IsDigit(char) {
				currentState = number
				numOfRepeat, _ = strconv.Atoi(string(char))
				repStr, err := repeatRune(runeArray[i-1], numOfRepeat)
				if err != nil {
					// Если функция repeatRune возвращает ошибку, удаляем последний символ из результата.
					res := result.String()
					result.Reset()
					result.WriteString(res[0 : len(res)-1])
				}
				result.WriteString(repStr)
			} else if char == '\\' {
				currentState = escape
			} else {
				result.WriteRune(char)
				currentState = start
			}
		case number:
			if unicode.IsDigit(char) {
				// Если после цифры идет ещё одна цифра, возвращаем ошибку "недопустимой строки".
				return "", ErrInvalidString
			}
			result.WriteRune(char)
			currentState = start
		case escape:
			// Если предыдущий символ - обратный слэш, добавляем текущий в результат.
			result.WriteRune(char)
			currentState = start
		}

	}

	return result.String(), nil

}

// repeatRune - Функция для повторения символа заданное количество раз.
func repeatRune(char rune, numOfRepeat int) (string, error) {
	if numOfRepeat > 0 {
		// Повторяем символ заданное количество раз и возвращаем строку.
		return strings.Repeat(string(char), numOfRepeat-1), nil
	}
	if numOfRepeat == 0 {
		// Если количество повторений равно нулю, возвращаем ошибку "нулевое повторение".
		return "", errors.New("zero repeat")
	}
	// Возвращаем пустую строку для отрицательных значений numOfRepeat.
	return "", errors.New("negative repeat")
}
