package isValidSudoku

import (
	"fmt"
	"testing"
)

func isValidSudoku(board [][]byte) bool {
	// 行列九宫的数据存在
	var rowArray, colArray [9][9]int
	var boxArray [9][9]int
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			// 读出数据
			cr := board[r][c]
			if cr == '.' {
				continue
			}
			index := cr - '1'
			// rowH
			rowArray[r][index]++
			// colH
			colArray[c][index]++
			// boxH 也可以将 boxArray 定义为三维数组
			// colIndex / 3 = 横向第N个box
			// rowIndex / 3 = 0~1~2 标识第几排的Box * 3 可以偏移 到第二排 第三排
			boxArray[c/3+(r/3)*3][index]++
			// 行列九宫存在数据
			if rowArray[r][index] > 1 || colArray[c][index] > 1 || boxArray[c/3+(r/3)*3][index] > 1 {
				return false
			}
		}
	}
	return true
}

func Test36(t *testing.T) {
	strArray := []string{
		"53..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	}
	board := make([][]byte, 0)
	for _, s := range strArray {
		board = append(board, []byte(s))
	}
	fmt.Println(isValidSudoku(board))

	strArray = []string{
		"83..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	}
	board = make([][]byte, 0)
	for _, s := range strArray {
		board = append(board, []byte(s))
	}
	fmt.Println(isValidSudoku(board))
}
