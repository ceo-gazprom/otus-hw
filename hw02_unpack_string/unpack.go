package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result strings.Builder
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		current := runes[i]

		if IsDigit(current) {
			return "", ErrInvalidString
		}

		if current == '\\' {
			if i+1 >= len(runes) || (!IsDigit(runes[i+1]) && runes[i+1] != '\\') {
				return "", ErrInvalidString
			}

			current = runes[i+1]
			i++
		}

		repeat := 1
		if i+1 < len(runes) && IsDigit(runes[i+1]) {
			repeat = int(runes[i+1] - '0')
			i++
		}

		result.WriteString(strings.Repeat(string(current), repeat))
	}
	return result.String(), nil
}

func IsDigit(num rune) bool {
	return num >= '0' && num <= '9'
}
