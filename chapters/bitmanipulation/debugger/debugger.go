package debugger

/*
Debugger: Explain what the following code does: ( ( n & ( n-1)) == 0).
*/

func IsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func NextPowerOfTwo(n int) int {
	if n == 0 {
		return 1
	}
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	return n + 1
}

func PrevPowerOfTwo(n int) int {
	if n == 0 {
		return 0
	}
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	return n ^ (n >> 1)
}
