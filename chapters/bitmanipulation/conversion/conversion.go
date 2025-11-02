package conversion

/*
Conversion: Write a function to determine the number of bits you would need to flip to convert
integer A to integer B.
EXAMPLE
Input: 29 (or: 11101), 15 (or: 01111)
Output: 2
*/

func Conversion(a, b int) int {
	diff := a ^ b
	count := 0
	for diff != 0 {
		diff &= diff - 1
		count++
	}
	return count
}
