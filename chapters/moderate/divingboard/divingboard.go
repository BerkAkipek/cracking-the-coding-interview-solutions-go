package divingboard

/*
Diving Board: You are building a diving board by placing a bunch of planks of wood end-to-end.
There are two types of planks, one of length shorter and one of length longer. You must use
exactly K planks of wood. Write a method to generate all possible lengths for the diving board.
*/

func DivingBoard(short, long, k int) []int {
	if k == 0 {
		return nil
	}
	if short == long {
		return []int{short * k}
	}
	if short > long {
		return nil
	}
	maxLength := long * k
	lengths := []int{maxLength}
	diff := long - short
	for range k {
		maxLength -= diff
		lengths = append(lengths, maxLength)
	}

	return lengths
}
