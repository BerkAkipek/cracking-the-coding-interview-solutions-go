package urlify

import "strings"

/*
URLify: Write a method to replace all spaces in a string with '%20'.
You may assume that the string has sufficient space at the end to hold the additional characters, and that you are given the "true" length of the string.
EXAMPLE
Input: "Mr John Smith ", 13
Output: "Mr%20John%20Smith"
*/

// Urlify is a method that replaces every space in a given string with '%20'
func Urlify(str string, trueLength int) string {
	str = str[:trueLength]

	var builder strings.Builder
	for _, ch := range str {
		if ch == ' ' {
			builder.WriteString("%20")
		} else {
			builder.WriteRune(ch)
		}
	}

	return builder.String()
}
