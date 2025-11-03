package sortedmerge

/*
Sorted Merge: You are given two sorted arrays, A and B, where A has a large enough buffer at the
end to hold B. Write a method to merge B into A in sorted order.
*/

func SortedMerge(a, b []int) []int {
	i, j, k := len(a)-len(b)-1, len(b)-1, len(a)-1
	for j >= 0 {
		if i >= 0 && a[i] > b[j] {
			a[k] = a[i]
			i--
		} else {
			a[k] = b[j]
			j--
		}
		k--
	}
	return a
}
