package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const backSlash = `\`

func Unpack(s string) (string, error) {
	if formatIsOk := checkStringFormat(s); !formatIsOk {
		return "", ErrInvalidString
	}

	var result strings.Builder
	var prevChar rune
	var prevIsBackSlash bool
	for index, value := range s {
		if prevIsBackSlash {
			result.WriteRune(prevChar)
			prevChar = value
			prevIsBackSlash = false
			continue
		}
		if string(value) == backSlash {
			prevIsBackSlash = true
			continue
		}
		if unicode.IsDigit(value) {
			count, _ := strconv.Atoi(string(value))
			result.WriteString(strings.Repeat(string(prevChar), count))
			prevChar = 0
		} else {
			if index > 0 && prevChar != 0 {
				result.WriteRune(prevChar)
			}
			prevChar = value
		}
	}
	if prevChar != 0 {
		result.WriteRune(prevChar)
	}
	return result.String(), nil
}

func checkStringFormat(s string) bool {
	isMatched, _ := regexp.Match(`^(([a-zA-Z]|\n)+\d?|(\\+\d+))*$`, []byte(s))
	return isMatched
}
