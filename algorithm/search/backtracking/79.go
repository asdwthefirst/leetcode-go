/**
  @author: jiangxi
  @since: 2022/12/30
  @desc: //TODO
**/
package backtracking

//超时，，感觉跟官解一样，为什么超时了，，，尼玛，把打印行删掉又不超时了
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	visited := make([][]bool, len(board))
	for i2 := range visited {
		visited[i2] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				success := backtrack79(board, visited, i, j, word[1:])
				if success {
					return true
				}
			}

		}

	}
	return false

}

func backtrack79(board [][]byte, visited [][]bool, i, j int, word string) bool {
	visited[i][j] = true
	if len(word) == 0 {
		return true
	}
	posNode := [][]int{{i + 1, j}, {i - 1, j}, {i, j + 1}, {i, j - 1}}
	for _, node := range posNode {
		inX, inY := node[0], node[1]
		if inX >= 0 && inY >= 0 && inX < len(board) && inY < len(board[0]) && visited[inX][inY] == false && board[inX][inY] == word[0] {
			success := backtrack79(board, visited, inX, inY, word[1:])
			if success {
				return success
			}
		}
	}
	visited[i][j] = false //关键，
	return false
}

//type pair struct{ x, y int }
//
//var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
//
//func exist(board [][]byte, word string) bool {
//	h, w := len(board), len(board[0])
//	vis := make([][]bool, h)
//	for i := range vis {
//		vis[i] = make([]bool, w)
//	}
//	var check func(i, j, k int) bool
//	check = func(i, j, k int) bool {
//		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
//			return false
//		}
//		if k == len(word)-1 { // 单词存在于网格中
//			return true
//		}
//		vis[i][j] = true
//		defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
//		for _, dir := range directions {
//			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
//				if check(newI, newJ, k+1) {
//					return true
//				}
//			}
//		}
//		return false
//	}
//	for i, row := range board {
//		for j := range row {
//			if check(i, j, 0) {
//				return true
//			}
//		}
//	}
//	return false
//}
