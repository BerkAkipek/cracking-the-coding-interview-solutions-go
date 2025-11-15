package subsort

import "math"

/*
Sub Sort: Given an array of integers, write a method to find indices m and n such that if you sorted
elements m through n, the entire array would be sorted. Minimize n - m (that is, find the smallest
such sequence).
EXAMPLE
Input: lj 2, 4j 7, 10, 11, 1, 12, 6, 7, 16, 18, 19
Output: (3 , 9)
*/

func subsort(arr []int) (int, int) {
	n := len(arr)
	if n <= 1 {
		return -1, -1
	}

	maxSeen := math.MinInt
	right := -1

	for i := 0; i < n; i++ {
		if arr[i] < maxSeen {
			right = i
		} else {
			maxSeen = arr[i]
		}
	}

	if right == -1 {
		return -1, -1
	}

	minSeen := math.MaxInt
	left := -1

	for i := n - 1; i >= 0; i-- {
		if arr[i] > minSeen {
			left = i
		} else {
			minSeen = arr[i]
		}
	}

	return left, right
}
