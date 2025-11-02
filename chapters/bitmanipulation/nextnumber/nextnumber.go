package nextnumber

/*
Next Number: Given a positive integer, print the next smallest and the next largest number that
have the same number of 1 bits in their binary representation.
*/

func NextNumber(num int) (int, int) {
	if num <= 0 {
		return 0, 0
	}

	c := num
	c0, c1 := 0, 0
	for (c&1) == 0 && c != 0 {
		c0++
		c >>= 1
	}
	for (c & 1) == 1 {
		c1++
		c >>= 1
	}

	p := c0 + c1

	nextBig := num | (1 << p)
	nextBig &= ^((1 << p) - 1)
	nextBig |= (1 << (c1 - 1)) - 1

	return nextBig, nextSmaller(num)
}

func nextSmaller(num int) int {
	c := num
	c1, c0 := 0, 0

	for (c & 1) == 1 {
		c1++
		c >>= 1
	}
	if c == 0 {
		return 0
	}
	for (c&1) == 0 && c != 0 {
		c0++
		c >>= 1
	}

	p := c0 + c1
	num &^= (1 << (p + 1)) - 1
	mask := (1 << (c1 + 1)) - 1
	num |= mask << (c0 - 1)
	return num
}
