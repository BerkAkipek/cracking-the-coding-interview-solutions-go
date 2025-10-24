package rotatematrix

/*
Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes,
write a method to rotate the image by 90 degrees.
Can you do this in place?
*/

func RotateMatrix(mtr [][]int) {
	n := len(mtr)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			temp := mtr[first][i]
			mtr[first][i] = mtr[last-offset][first]
			mtr[last-offset][first] = mtr[last][last-offset]
			mtr[last][last-offset] = mtr[i][last]
			mtr[i][last] = temp
		}
	}
}
