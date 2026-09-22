package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	// Делим поп робелам, переносам и тдж
	parts := strings.Fields(text)
	words := make(map[string]int)

	// Получаем кол-во использований для каждого слова
	for _, word := range parts {
		words[word]++
	}

	// Делаем слайс для сортировки
	result := make([]string, 0, len(words))

	for word := range words {
		result = append(result, word)
	}

	sort.Slice(result, func(i, j int) bool {
		// если потворения одинаковы
		if words[result[i]] == words[result[j]] {
			// по алфавиту
			return result[i] < result[j]
		}

		return words[result[i]] > words[result[j]]
	})

	// Возвращаем 10
	if len(result) > 10 {
		result = result[:10]
	}

	return result
}
