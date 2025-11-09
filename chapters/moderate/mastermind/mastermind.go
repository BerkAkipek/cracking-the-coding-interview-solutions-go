package mastermind

import "math/rand"

/*
Master Mind: The Game of Master Mind is played as follows:
The computer has four slots, and each slot will contain a ball that is red (ft), yellow (V), green (G) or
blue (B). For example, the computer might have RGGB (Slot #1 is red, Slots #2 and #3 are green. Slot
#4 is blue).
You, the user, are trying to guess the solution. You might, for example, guess YRGB.
When you guess the correct color for the correct slot, you get a "hit." If you guess a color that exists
but is in the wrong slot, you get a "pseudo-hit." Mote that a slot that is a hit can never count as a
pseudo-hit.
For example, if the actual solution is RGBY and you guess GGRR, you have one hit and one pseudo-hit.
Write a method that, given a guess and a solution, returns the number of hits and pseudo-hits.
*/

func Setup() ([]string, []string) {
	space := []string{"R", "G", "B", "Y"}
	computer := []string{}
	player := []string{}

	for range 4 {
		computer = append(computer, space[rand.Intn(len(space))])
		player = append(player, space[rand.Intn(len(space))])
	}

	return computer, player
}

func MasterMind(computer, player []string) []string {
	result := []string{}

	for i := range len(computer) {
		if player[i] == computer[i] {
			result = append(result, "hit")
			continue
		} else {
			if checkPseudoHit(i, player[i], computer) {
				result = append(result, "pseudo-hit")
			}
		}
	}

	return result
}

func checkPseudoHit(ind int, guess string, elements []string) bool {
	for i, e := range elements {
		if e == guess {
			if i != ind {
				return true
			}
		}
	}
	return false
}
