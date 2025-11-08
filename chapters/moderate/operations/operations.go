package operations

/*
Operations: implement multiply, subtract, and divide for integers.
The results are integers. Use only the add operator.
*/

func Subtraction(x, y int) int {
	return x + negate(y)
}

func Multiplication(x, y int) int {
	sum := 0
	a, b := abs(x), abs(y)

	for range b {
		sum += a
	}

	if (x < 0 && y > 0) || (x > 0 && y < 0) {
		return negate(sum)
	}
	return sum
}

func Division(x, y int) int {
	if y == 0 {
		panic("division by zero")
	}

	count := 0
	a, b := abs(x), abs(y)

	for a >= b {
		a += negate(b)
		count++
	}

	if (x < 0 && y > 0) || (x > 0 && y < 0) {
		return negate(count)
	}
	return count
}

func abs(x int) int {
	if x < 0 {
		return negate(x)
	}
	return x
}

func negate(n int) int {
	neg := 0
	delta := 1
	if n > 0 {
		delta = -1
	}
	for n != 0 {
		n += delta
		neg += delta
	}
	return neg
}
