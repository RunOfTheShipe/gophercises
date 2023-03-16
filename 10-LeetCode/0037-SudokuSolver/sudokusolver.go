package sudokusolver

// puzzle 37
func solveSudoku(board [][]byte) {
	solvePuzzle(board)
}

func solvePuzzle(board [][]byte) bool {

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {

			// check to see if there is already an item in the
			// solution; if not, then try to guess the number
			if board[r][c] == '.' {

				for g := byte('1'); g <= '9'; g++ {

					// if it's a valid guess, try to place the number
					// and try to solve the puzzle
					if validGuess(r, c, g, board) {
						board[r][c] = g

						if solvePuzzle(board) {
							// puzzle solved - stop guessing
							return true
						} else {
							// problem wasn't solved, reset back to default
							board[r][c] = '.'
						}
					}
				}

				// tried all 9 numbers and none of them worked, return
				// false to backtrack and try a different number in the
				// previous spot
				return false
			}
		}
	}

	// only this this when all squares are filled
	return true
}

func validGuess(row int, col int, guess byte, board [][]byte) bool {
	// check the row for the same value
	for i := 0; i < 9; i++ {
		if board[i][col] == guess {
			return false
		}
	}

	// check the column for the same value
	for i := 0; i < 9; i++ {
		if board[row][i] == guess {
			return false
		}
	}

	// check the inner grid for the same value
	rowStart := (row / 3) * 3
	colStart := (col / 3) * 3
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			if board[i][j] == guess {
				return false
			}
		}
	}

	return true
}

// puzzle 36
func isValidSudoku(board [][]byte) bool {

	var temp byte = 0

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			// skip anything that isn't a '.'
			if board[i][j] != '.' {
				temp = board[i][j]
				board[i][j] = '.'
				if !validGuess(i, j, temp, board) {
					return false
				}
				board[i][j] = temp
			}
		}
	}
	return true
}
