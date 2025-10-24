package stringcompression

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

/*
Implement a method to perform basic string compression using the counts of consecutive repeated characters.
If the "compressed" string would not become smaller than the original string, your method should return the original string.

You may assume the input string consists only of uppercase and lowercase English letters (A–Z, a–z).
*/

// StringCompression performs basic string compression using the counts of consecutive repeated characters.
// If the "compressed" string would not become smaller than the original string, it returns the origiinal string.
func StringCompression(str string) string {
	if len(str) == 0 {
		return str
	}

	var builder strings.Builder
	count := 0
	currentChar, _ := utf8.DecodeRuneInString(str)
	for _, char := range str {
		if currentChar != char {
			builder.WriteString(string(currentChar))
			builder.WriteString(strconv.Itoa(count))
			currentChar = char
			count = 1
		} else {
			count++
		}
	}

	builder.WriteString(string(currentChar))
	builder.WriteString(strconv.Itoa(count))

	result := builder.String()
	if len(result) < len(str) {
		return result
	}
	return str
}
