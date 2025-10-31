package eggdropproblem

import "math/rand"

/*
The Egg Drop Problem: There is a building of 100 floors. If an egg drops from the Nth floor or
above, it will break. If it's dropped from any floor below, it will not break. You're given two eggs. Find
N, while minimizing the number of drops for the worst case
*/

// FindThreshold returns the smallest floor N such that breaks(N)==true.
// It uses the optimal triangular-step strategy with 2 eggs.
func FindThreshold(numFloors int, breaks func(int) bool) int {
	step := 0
	for step*(step+1)/2 < numFloors {
		step++
	}

	current := step
	prev := 0

	// Phase 1: use egg #1 with decreasing steps.
	for current <= numFloors {
		if breaks(current) { // egg #1 broke here
			break
		}
		step--
		prev = current
		current += step
	}

	// Phase 2: linear search with egg #2 from (prev+1) up to current.
	upper := current
	if upper > numFloors {
		upper = numFloors
	}
	for f := prev + 1; f <= upper; f++ {
		if breaks(f) {
			return f
		}
	}

	// If it never broke, threshold is above the building; define as top.
	return numFloors
}

func TriangularSearch(numFloors int) int {
	N := rand.Intn(numFloors) + 1

	return FindThreshold(numFloors, func(floor int) bool {
		return floor >= N
	})
}
