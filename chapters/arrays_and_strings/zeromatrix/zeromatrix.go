package zeromatrix

/*
Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column are set to 0.
*/

// ZeroMatrix implements an algorithm such that if an element in an MxN matrix is 0 its entire row and column are set to 0.
func ZeroMatrix(mtr [][]int) [][]int {
	res := make([][]int, len(mtr))
	for i := range mtr {
		res[i] = make([]int, len(mtr[i]))
		copy(res[i], mtr[i])
	}

	rows := make([]bool, len(mtr))
	cols := make([]bool, len(mtr[0]))
	for i := range mtr {
		for j := 0; j < len(mtr[i]); j++ {
			if mtr[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i := range mtr {
		for j := 0; j < len(mtr[i]); j++ {
			if rows[i] || cols[j] {
				res[i][j] = 0
			}
		}
	}

	return res
}
