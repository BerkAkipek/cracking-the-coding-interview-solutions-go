package dominos

/*
Dominos: There is an 8x8 chessboard in which two diagonally opposite corners have been cut off.
You are given 31 dominos, and a single domino can cover exactly two squares. Can you use the 31
dominos to cover the entire board? Prove your answer (by providing an example or showing why
it's impossible).
*/
const N = 8

type Cell struct{ R, C int }

// 0 = black, 1 = white
func color(r, c int) int { return (r + c) & 1 }

// CanTile checks color parity after removing cells.
func CanTile(removed []Cell) (possible bool, blackCount, whiteCount int) {
	rm := make(map[Cell]bool, len(removed))
	for _, c := range removed {
		rm[c] = true
	}
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if rm[Cell{r, c}] {
				continue
			}
			if color(r, c) == 0 {
				blackCount++
			} else {
				whiteCount++
			}
		}
	}
	return blackCount == whiteCount, blackCount, whiteCount
}
