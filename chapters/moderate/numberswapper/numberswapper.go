package numberswapper

/*
Number Swapper: Write a function to swap a number in place (that is, without temporary variables).
*/

func Swapper(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}
