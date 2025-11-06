package wordfrequencies

import "strings"

/*
Word Frequencies: Design a method to find the frequency of occurrences of any given word in a book.
What if we were running this algorithm multiple times?
*/

func WordFrequencies(book []string) map[string]int {
	result := make(map[string]int)
	punctuations := map[rune]bool{
		'.': true, ',': true, '?': true, '!': true, ':': true, ';': true,
		'-': true, '—': true, '(': true, ')': true, '[': true, ']': true,
		'{': true, '}': true, '\'': true, '"': true, '…': true,
		'“': true, '”': true, '‘': true, '’': true, // added curly quotes
	}

	for _, line := range book {
		words := strings.Fields(line)
		for _, word := range words {
			word = strings.ToLower(word)
			word = strings.TrimSuffix(word, "...")
			word = trimPunctuation(word, punctuations)
			if word == "" {
				continue
			}
			result[word]++
		}
	}
	return result
}

func trimPunctuation(s string, punct map[rune]bool) string {
	runes := []rune(s)
	start, end := 0, len(runes)-1
	for start <= end && punct[runes[start]] {
		start++
	}
	for end >= start && punct[runes[end]] {
		end--
	}
	return string(runes[start : end+1])
}
