package numbermax

/*
Number Max: Write a method that finds the maximum of two numbers. You should not use if-else
or any other comparison operator.
*/

func CompareNums(x, y int) int {
	sign_x := (uint64(x) >> 63) & 1
	sign_y := (uint64(y) >> 63) & 1

	diff_sign := sign_x ^ sign_y
	signDiff := (uint64((x - y)) >> 63) & 1
	useDiff := ^diff_sign & 1

	finalSign := diff_sign*sign_x + useDiff*signDiff
	mask := int64(finalSign * ^uint64(0))

	return int((int64(x) & ^mask) | (int64(y) & mask))
}
