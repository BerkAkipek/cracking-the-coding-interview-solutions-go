package hundredlockers

import "sort"

/*
100 Lockers: There are 100 closed lockers in a hallway. A man begins by opening all 100 lockers.
Next, he closes every second locker. Then, on his third pass, he toggles every third locker (closes it if
it is open or opens it if it is closed). This process continues for 100 passes, such that on each pass i,
the man toggles every ith locker. After his 100th pass in the hallway, in which he toggles only locker
how many lockers are open
*/

func Lockers(num int) []int {
	lockers := make(map[int]bool)
	for i := range num {
		lockers[i+1] = false
	}

	for i := 1; i <= num; i++ {
		for key := range lockers {
			if key%i == 0 {
				lockers[key] = !lockers[key]
			}
		}
	}

	result := []int{}
	for key, val := range lockers {
		if val {
			result = append(result, key)
		}
	}
	sort.Ints(result)
	return result
}
