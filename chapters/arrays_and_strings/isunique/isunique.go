package isunique

/*
Is Unique: Implement an algorithm to determine if a string has all unique characters.
What if you cannot use additional data structures?
*/

// IsUnique returns true if all characters of given string is unique
func IsUnique(str string) bool {
	seen := make(map[rune]bool)

	for _, ch := range str {
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}
