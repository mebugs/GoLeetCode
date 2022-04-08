package setZeroes

import (
	"testing"
)

func setZeroes(matrix [][]int) {
	rL := len(matrix)
	cL := len(matrix[0])
	// 标识变零
	rowMap := make(map[int]bool, rL)
	colMap := make(map[int]bool, cL)
	// 开始遍历标志
	for r, row := range matrix {
		for c, col := range row {
			if col == 0 {
				rowMap[r] = true
				colMap[c] = true
			}
		}
	}
	// 开始处理数据
	for r := 0; r < rL; r++ {
		for c := 0; c < cL; c++ {
			if rowMap[r] || colMap[c] {
				matrix[r][c] = 0
			}
		}
	}
}

func Test73(t *testing.T) {
	setZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})
	setZeroes([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}})
}
