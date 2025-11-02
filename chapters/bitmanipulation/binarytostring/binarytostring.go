package binarytostring

/*
Binary to String: Given a real number between O and 1 (e.g., 0.72) that is passed in as a double, print
the binary representation. If the number cannot be represented accurately in binary with at most 32
characters, print "ERROR:'
*/

func BinaryRepresentation(num float64) string {
	if num <= 0 || num >= 1 {
		return "ERROR"
	}

	result := "."
	for num > 0 {
		if len(result) >= 32 {
			return "ERROR"
		}
		num *= 2
		if num >= 1 {
			result += "1"
			num -= 1
		} else {
			result += "0"
		}
	}
	return "0" + result
}
