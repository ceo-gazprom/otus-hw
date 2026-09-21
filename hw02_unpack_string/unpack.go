package main

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func Unpack(s string) (string, error) {
	runes := []rune(s)

	var result strings.Builder

	for i := 0; i < len(runes); i++ {
		current := runes[i]

		// если встретили слэш
		if current == '\\' {
			// после слэша должен быть следующий символ
			if i+1 >= len(runes) {
				return "", ErrInvalidString
			}

			next := runes[i+1]

			// экранировать можно только цифру или слэш
			if !isDigit(next) && next != '\\' {
				return "", ErrInvalidString
			}

			current = next

			// Пропускаем обработанную руну после слэша
			i++
		} else if isDigit(current) {
			return "", ErrInvalidString
		}

		count := 1

		// если после руны стоит цифра,
		// будем записывать повтор
		if i+1 < len(runes) && isDigit(runes[i+1]) {
			count = int(runes[i+1] - '0')
			i++

			// есил 2 цифры подряд, то запрещено
			if i+1 < len(runes) && isDigit(runes[i+1]) {
				return "", ErrInvalidString
			}
		}

		// Записываем n раз символ текущий (если выше была цифра)
		result.WriteString(strings.Repeat(string(current), count))
	}

	return result.String(), nil
}
