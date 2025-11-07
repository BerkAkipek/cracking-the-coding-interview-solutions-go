package tictac

/*
TicTac Win: Design an algorithm to figure out if someone has won a game of tic-tac-toe.
*/

func WhoWins(board [][]int) int {
	diagonal := checkDiagonals(board)
	if diagonal != -1 {
		return diagonal
	}

	row := checkRows(board)
	if row != -1 {
		return row
	}

	column := checkColums(board)
	if column != -1 {
		return column
	}

	return -1
}

func checkDiagonals(board [][]int) int {
	element := board[0][0]
	if element == board[1][1] && element == board[2][2] {
		return element
	}

	element = board[0][2]
	if element == board[1][1] && element == board[2][0] {
		return element
	}

	return -1
}

func checkRows(board [][]int) int {
	for i := range 3 {
		element := board[i][0]
		if element == board[i][1] && element == board[i][2] {
			return element
		}
	}
	return -1
}

func checkColums(board [][]int) int {
	for i := range 3 {
		element := board[0][i]
		if element == board[1][i] && element == board[2][i] {
			return element
		}
	}
	return -1
}
