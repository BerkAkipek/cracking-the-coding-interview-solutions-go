package groupanagrams

import "slices"

/*
Group Anagrams: Write a method to sort an array of strings so that all the anagrams are next to each other.
*/

func GroupAnagrams(arr []string) []string {
	groups := make(map[string][]string)

	for i := range len(arr) {
		key := sortString(arr[i])
		groups[key] = append(groups[key], arr[i])
	}

	var result []string
	for _, group := range groups {
		result = append(result, group...)
	}
	return result
}

func sortString(s string) string {
	r := []rune(s)
	slices.Sort(r)
	return string(r)
}
