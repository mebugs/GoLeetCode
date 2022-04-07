package matrixReshape

import "testing"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	// 判定异常返回
	rO := len(mat)
	cO := len(mat[0])
	if rO*cO != r*c || (r == 0 || c == 0) {
		return mat
	}
	// 初始化新矩阵
	nMat := make([][]int, r)
	for index := range nMat {
		nMat[index] = make([]int, c)
	}
	// 遍历组装新矩阵
	for ni := 0; ni < rO*cO; ni++ {
		// ni除以c = 行坐标
		// ni取余c = 纵坐标
		// ni除以cO = 旧数据对应的行坐标
		// ni取余cO = 旧数据对应的纵坐标
		nMat[ni/c][ni%c] = mat[ni/cO][ni%cO]
	}
	return nMat
}

func Test566(t *testing.T) {
	t.Log(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
	t.Log(matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4))
	t.Log(matrixReshape([][]int{{6, 2}, {3, 9}, {8, 7}}, 2, 3))
}
