package oneaway

// OneAway returns true if target can be created with one edit
// An edit defined as inserting, removing, or replacing a single character.
func OneAway(input, target string) bool {
	// Base case return false
	if abs(len(input)-len(target)) > 1 {
		return false
	}
	i, j, diff := 0, 0, 0
	for i < len(input) && j < len(target) {
		if input[i] != target[j] {
			diff++
			if diff > 1 {
				return false
			}
			// Meaty part
			if len(input) > len(target) {
				i++
				continue
			}
			if len(target) > len(input) {
				j++
				continue
			}
		}
		i++
		j++
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
