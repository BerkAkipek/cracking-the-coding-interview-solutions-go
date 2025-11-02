package insertion

/*
Insertion: You are given two 32-bit numbers, N and M, and two bit positions, i and
j. Write a method to insert M into N such that M starts at bit j and ends at bit i. You
can assume that the bits j through i have enough space to fit all of M. That is, if
M = 10011, you can assume that there are at least 5 bits between j and i. You would not, for
example, have j = 3 and i = 2, because M could not fully fit between bit 3 and bit 2.
EXAMPLE
Input: N 10000000000, M
Output: N = 10001001100
*/

func Insertion(n, m, i, j int32) int32 {
	if i > j || i < 0 || j >= 32 {
		return 0
	}

	var allOnes int32 = ^0
	left := allOnes << (j + 1)
	var right int32 = (1 << i) - 1
	mask := left | right

	nCLeared := n & mask
	mShifted := m << i
	return nCLeared | mShifted
}
