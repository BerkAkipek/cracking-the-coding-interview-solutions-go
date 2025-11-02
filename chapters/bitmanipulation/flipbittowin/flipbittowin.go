package flipbittowin

/*
Flip Bit to Win: You have an integer and you can flip exactly one bit from a 0 to a 1. Write code to
find the length of the longest sequence of ls you could create.
EXAMPLE
Input: 1775
Output: 8
*/

func FlipToWin(num int) int {
	if num == 0 {
		return 1
	}
	if num < 0 {
		return 0
	}
	if num == (1<<32 - 1) {
		return 32
	}

	currentLen, prevLen, maxLen := 0, 0, 0
	zeroSeen := false

	for num != 0 {
		if num&1 == 1 {
			currentLen++
		} else {
			if !zeroSeen {
				prevLen = currentLen
				zeroSeen = true
			} else {
				prevLen = 0
			}
			currentLen = 0
		}

		maxLen = max(maxLen, prevLen+currentLen+1)
		num >>= 1

		if num&1 == 1 {
			zeroSeen = false
		}
	}

	return maxLen
}
