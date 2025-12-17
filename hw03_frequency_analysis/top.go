package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	// Place your code here.
	if len(strings.TrimSpace(s)) == 0 {
		return nil
	}

	words := strings.Fields(s)

	if len(words) == 0 {
		return nil
	}

	count := make(map[string]int)

	for _, word := range words {
		if word != "" {
			count[word]++
		}
	}

	type wordValue struct {
		word  string
		value int
	}

	var result []wordValue

	for word, value := range count {
		result = append(result, wordValue{word, value})
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].value != result[j].value {
			return result[i].value > result[j].value
		}
		return result[i].word < result[j].word
	})

	ten := len(result)

	if ten > 10 {
		ten = 10
	}

	totalResult := make([]string, ten)

	for i := 0; i < ten; i++ {
		totalResult[i] = result[i].word
	}

	return totalResult
}
