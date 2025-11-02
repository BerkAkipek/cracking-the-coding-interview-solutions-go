package poison

import "math/rand"

/*
Poison: You have 1000 bottles of soda, and exactly one is poisoned. You have 10 test strips which
can be used to detect poison. A single drop of poison will turn the test strip positive permanently.
You can put any number of drops on a test strip at once and you can reuse a test strip as many times
as you'd like (as long as the results are negative). However, you can only run tests once per day and
it takes seven days to return a result. How would you figure out the poisoned bottle in as few days
as possible?
FOLLOW UP
Write code to simulate your approach
*/

type bottle struct {
	id         int
	isPoisoned bool
}

func Setup() ([]*bottle, []bool) {
	bottles := []*bottle{}
	strips := []bool{}
	for i := range 1000 {
		bottles = append(bottles, &bottle{id: i, isPoisoned: false})
	}
	for range 10 {
		strips = append(strips, false)
	}

	poisoned := bottles[rand.Intn(len(bottles))]
	poisoned.isPoisoned = true

	return bottles, strips
}

func EncodeStrips(bottles []*bottle, strips []bool) {
	for _, b := range bottles {
		for bit := 0; bit < 10; bit++ {
			if ((b.id>>bit)&1) == 1 && b.isPoisoned {
				strips[bit] = true
			}
		}
	}
}

func DecodeStrips(strips []bool) int {
	result := 0
	for i := 0; i < len(strips); i++ {
		if strips[i] {
			result |= 1 << i
		}
	}
	return result
}
