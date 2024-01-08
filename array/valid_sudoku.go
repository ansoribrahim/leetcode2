package array

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = map[byte]bool{}
		cols[i] = map[byte]bool{}
		boxes[i] = map[byte]bool{}
	}

	for i := 0; i < 9; i++ {
		fmt.Println()
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			// box will be count from left to right, start from box 0 to 2, 3 to 5, 6 to 8
			// (i/3)*3 will indicate row of box if row not more than 3, it will be in row 0,
			// if i less than 3, it will be in row 0. ex : i = 2, then result will be 0. i = 3, then result will be 1
			// j/3 this will indicate column of box, divided by 3, because j can be 9, and there are 3 boxes there
			box := (i/3)*3 + j/3

			if !isValid(board[i][j], rows[i]) ||
				!isValid(board[i][j], cols[j]) ||
				!isValid(board[i][j], boxes[box]) {
				return false
			}
		}
	}

	return true
}

func isValid(num byte, set map[byte]bool) bool {
	if set[num] {
		return false
	}
	set[num] = true
	return true
}
