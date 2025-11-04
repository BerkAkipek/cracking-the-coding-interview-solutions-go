package sparsesearch

/*
Sparse Search: Given a sorted array of strings that is interspersed with empty strings, write a
method to find the location of a given string.
EXAMPLE
Input: ball,{"at" , "" , "" , "" , "ball" , "" , "car" , "" , "" , "dad" , "" ,
any
Output: 4
*/

func SparseSearch(arr []string, target string) int {
	if len(arr) == 0 || target == "" {
		return -1
	}

	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == "" {
			left, right := mid-1, mid+1
			for {
				if left < low && right > high {
					return -1
				}
				if right <= high && arr[right] != "" {
					mid = right
					break
				}
				if left >= low && arr[left] != "" {
					mid = left
					break
				}
				right++
				left--
			}
		}

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
