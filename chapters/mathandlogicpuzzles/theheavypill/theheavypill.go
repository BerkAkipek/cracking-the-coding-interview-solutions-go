package theheavypill

import (
	"fmt"
	"math/rand"
)

func sumUntilN(n int) int {
	return (n * (n + 1)) / 2
}

func setup() [][]float64 {
	bottles := [][]float64{}

	// 19 normal bottles
	for range 19 {
		carry := []float64{}
		for range 20 {
			carry = append(carry, 1.0)
		}
		bottles = append(bottles, carry)
	}

	// 1 special bottle
	special := []float64{}
	for range 20 {
		special = append(special, 1.1)
	}
	bottles = append(bottles, special)

	return bottles
}

// Fisher–Yates Shuffle
func shuffle(bottles [][]float64) {
	for i := len(bottles) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		bottles[i], bottles[j] = bottles[j], bottles[i]
	}
}

func Game() {
	bottles := setup()
	shuffle(bottles)

	expected := float64(sumUntilN(len(bottles))) // expected weight if all pills = 1.0g
	realSum := 0.0

	// Take 1 pill from bottle 1, 2 from bottle 2, etc.
	for i := range bottles {
		count := i + 1
		for j := 0; j < count; j++ {
			realSum += bottles[i][j]
		}
	}

	diff := realSum - expected
	heavyIndex := int(diff*10) - 1 // because 0.1g extra per pill

	fmt.Printf("Expected: %.1f g\n", expected)
	fmt.Printf("Actual:   %.1f g\n", realSum)
	fmt.Printf("Heavy bottle is #%d\n", heavyIndex+1)
}
