package stringrotation

/*
String Rotation: Assume you have a method isSubstring which checks if one word is a substring of another.
Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one
call to i5Sub5tring (e.g., "waterbottle" is a rotation of"erbottlewat").

Pseudocode:
Isubstring will return a slice of strings contains all substringss

	Uses two pointer method. Both pointers start from left
	Append result with every substring

IsRotations will return true if the str2 is rotation of str1.

	Call IsSubstring method.
	count the length of each string.
	if len(str1) == count -> return true
	return false
*/
func IsSubstring(str1, str2 string) bool {
	i, j := 0, 0
	for i < len(str1) {
		if str1[i] == str2[j] {
			i++
			j++
			if j == len(str2) {
				return true
			}
		} else {
			i = i - j + 1
			j = 0
		}
	}
	return false
}

func IsRotation(str1, str2 string) bool {
	if len(str1) == 0 && len(str2) == 0 {
		return true
	}

	if len(str1) != len(str2) {
		return false
	}

	return IsSubstring(str1+str1, str2)
}
