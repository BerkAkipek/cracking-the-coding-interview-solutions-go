package sortedsearch

/*
Sorted Search, No Size: You are given an array-like data structure List y which lacks a size
method. It does, however, have an elementAt (i ) method that returns the element at index i in
0(1 ) time, if i is beyond the bounds of the data structure, it returns -1. (For this reason, the data
structure only supports positive integers.) Given a List y which contains sorted, positive integers,
find the index at which an element X occurs. If x occurs multiple times, you may return any index.
*/

func SortedSearch(arr []int, target int) int {
	if elementAt(arr, 0) == -1 {
		return -1
	}
	if elementAt(arr, 0) == target {
		return 0
	}

	i := 1
	for {
		val := elementAt(arr, i)
		if val == -1 || val >= target {
			break
		}
		i <<= 1
	}

	return binarySearch(arr, target, i/2, i)
}

func binarySearch(arr []int, target, low, high int) int {
	for low <= high {
		mid := (low + high) / 2
		val := elementAt(arr, mid)

		if val == -1 || val > target {
			high = mid - 1
		} else if val < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func elementAt(arr []int, i int) int {
	if i < 0 || i >= cap(arr) {
		return -1
	}
	return arr[i]
}
