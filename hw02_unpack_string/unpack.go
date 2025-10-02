package hw02unpackstring

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("некорректная строка")

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	var result []rune
	runes := []rune(s)
	i := 0

	for i < len(runes) {

		current := runes[i]
		i++

		if unicode.IsDigit(current) && len(result) == 0 {
			return "", ErrInvalidString
		}

		if i < len(runes) && unicode.IsDigit(runes[i]) {
			if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
				return "", ErrInvalidString
			}

			digit := runes[i]
			count := int(digit - '0')
			i++

			for j := 0; j < count; j++ {
				result = append(result, current)
			}
		} else {
			result = append(result, current)
		}
	}

	return string(result), nil
}
