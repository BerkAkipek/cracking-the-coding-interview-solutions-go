package palindromepermutation

import "unicode"

/*
Palindrome Permutation: Given a string, write a function to check if it is a permutation of
a palindrome. A palindrome is a word or phrase that is the same forwards and backwards. A
permutation is a rearrangement of letters. The palindrome does not need to be limited to just
dictionary words.
EXAMPLE
Input: Tact Coa
Output: True (permutations: "taco cat'; "atco eta·; etc.)
*/
func PalindromePermutation(str string) bool {
	counts := make(map[rune]int)
	for _, r := range str {
		if unicode.IsLetter(r) {
			r = unicode.ToLower(r)
			counts[r]++
		}
	}

	oddCount := 0
	for _, count := range counts {
		if count&1 != 0 {
			oddCount++
			if oddCount > 1 {
				return false
			}
		}
	}

	return true
}
