package englishint

import "strings"

/*
English Int: Given any integer, print an English phrase that describes the integer (e.g., "OneThousand, Two Hundred Thirty Four").
*/

func EnglishInt(n int) string {
	if n == 0 {
		return "zero"
	}

	result := ""
	if n < 0 {
		result = "negative " + EnglishInt(-n)
		return result
	}
	suffixes := []string{"", "thousand ", "million ", "billion ", "trillion ", "quadrillion ", "quintillion ", "sextillion ", "septillion ", "octillion "}

	i := 0
	for n != 0 {
		chunk := n % 1000
		if chunk != 0 {
			part := converChunk(chunk) + suffixes[i]
			result = part + result
		}
		n /= 1000
		i++
	}

	return strings.TrimSpace(result)
}

func converChunk(chunk int) string {
	digits := []string{"zero ", "one ", "two ", "three ", "four ", "five ", "six ", "seven ", "eight ", "nine "}
	tenth := []string{"", "", "twenty ", "thirty ", "forty ", "fifty ", "sixty ", "seventy ", "eighty ", "ninety "}
	teens := []string{"ten ", "eleven ", "twelve ", "thirteen ", "fourteen ", "fifteen ", "sixteen ", "seventeen ", "eighteen ", "nineteen "}

	if chunk == 0 {
		return ""
	}
	if chunk < 10 {
		return digits[chunk]
	}
	if chunk < 20 {
		return teens[chunk-10]
	}
	if chunk < 100 {
		tensPart := tenth[chunk/10]
		ones := chunk % 10
		if ones == 0 {
			return tensPart
		}
		return tensPart + digits[ones]
	}
	if chunk < 1000 {
		hundredsPart := digits[chunk/100] + "hundred "
		remainder := converChunk(chunk % 100)
		return hundredsPart + remainder
	}
	return ""
}
