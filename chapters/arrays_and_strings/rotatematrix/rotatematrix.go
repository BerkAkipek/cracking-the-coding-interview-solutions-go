package rotatematrix

/*
Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes,
write a method to rotate the image by 90 degrees.
Can you do this in place?
*/

var Matrix = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

func Transpose(mtr [][]int) {
	if mtr == nil {
		return
	}

	if len(mtr) == 0 {
		return
	}

	sanityLength := len(mtr[0])
	for i := 0; i < len(mtr); i++ {
		if sanityLength != len(mtr[i]) {
			return
		}
	}

	temp := 0
	for i := range mtr {
		for j := 0; j < len(mtr[i]); j++ {
			if j > i {
				temp = mtr[i][j]
				mtr[i][j] = mtr[j][i]
				mtr[j][i] = temp
			}
		}
	}
}

func RotateMatrix(mtr [][]int) {
	Transpose(mtr)
	if mtr == nil {
		return
	}

	if len(mtr) == 0 {
		return
	}

	rowInd := 0
	left, right := 0, len(mtr[0])-1
	for range len(mtr) {
		for left < right {
			temp := mtr[rowInd][left]
			mtr[rowInd][left] = mtr[rowInd][right]
			mtr[rowInd][right] = temp
			left++
			right--
		}
		rowInd++
		left, right = 0, len(mtr[0])-1
	}
}
