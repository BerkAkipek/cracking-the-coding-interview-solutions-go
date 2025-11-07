package factorialzeros

/*
Factorial Zeros: Write an algorithm which computes the number of trailing zeros in n factorial.
*/

func TrailingZeros(num int) int {
	if num < 5 {
		return 0

	}

	for num%5 != 0 {
		num--
	}

	zeros := 0
	for num > 0 {
		num /= 5
		zeros += num
	}

	return zeros
}
