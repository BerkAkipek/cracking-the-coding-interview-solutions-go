package checkpermutation

/*
Check Permutation: Given two strings, write a method to decide if one is a permutation of the other.
*/

// CheckPermutation takes two string as a parameters and returns true if the given string is a permutation of other
func CheckPermutation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	counts := make(map[rune]int)
	for _, ch := range str1 {
		counts[ch]++
	}
	for _, ch := range str2 {
		counts[ch]--
		if counts[ch] < 0 {
			return false
		}
	}

	return true
}
