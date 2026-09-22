package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	// финальная строка
	var result strings.Builder

	// например 		{input: "a4bc2d5e", expected: "aaaabccddddde"},
	// Разбиваем строку по символам
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		current := runes[i]
		repeat := 1

		// Проверка что цифры не удиут сразу или не 2 раза подряд
		if IsDigit(current) {
			return "", ErrInvalidString
		}

		// проверяем, есть ли слудующий символ
		if i+1 < len(runes) && IsDigit(runes[i+1]) {
			repeat = int(runes[i+1] - '0')
			// Пропускаем цифру в след иттерации
			i++
		}

		result.WriteString(strings.Repeat(string(current), repeat))
	}
	return result.String(), nil
}

func IsDigit(num rune) bool {
	return num >= '0' && num <= '9'
}
