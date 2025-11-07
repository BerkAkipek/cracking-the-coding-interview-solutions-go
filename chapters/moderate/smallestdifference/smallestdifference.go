package smallestdifference

import (
	"math"
	"slices"
)

/*
Smallest Difference: Given two arrays of integers, compute the pair of values (one value in each
array) with the smallest (non-negative) difference. Return the difference.
EXAMPLE
Input: {1, 3,1S, 11, 2}, (23,127, 235,19, 8)
Output: 3. That is, the pair (11,8).
*/

func SmallestDifference(arr1, arr2 []int) int {
	if len(arr1) == 0 || len(arr2) == 0 {
		return 0
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	i, j := 0, 0
	minDif := math.MaxInt

	for i < len(arr1) && j < len(arr2) {
		diff := abs(arr1[i] - arr2[j])

		if diff < minDif {
			minDif = diff
		}

		if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}

	return minDif
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
