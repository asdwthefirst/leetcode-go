/**
  @author: jiangxi
  @since: 2022/12/30
  @desc: //TODO
**/
package backtracking

func solveSudoku(board [][]byte) {
	var emptyPos [][2]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				emptyPos = append(emptyPos, [2]int{i, j})
			}
		}
	}

	var backtrack37 func(int) bool
	backtrack37 = func(index int) bool {
		x, y := emptyPos[index][0], emptyPos[index][1]

		for i := 1; i <= 9; i++ {
			if valid(board, x, y, i) {
				board[x][y] = byte('0' + i)
				if index == len(emptyPos)-1 {
					return true
				} else {
					success := backtrack37(index + 1)
					if success {
						return success
					}
					board[x][y] = '.'
				}
			}
		}
		return false
	}
	for i := 0; i < len(emptyPos); i++ {
		backtrack37(i)
	}
	return

}

func valid(board [][]byte, i, j int, num int) bool {
	for z := 0; z < 9; z++ {
		if board[i][z] == byte('0'+num) || board[z][j] == byte('0'+num) {
			return false
		}
	}
	areaX, areaY := i/3, j/3
	for i2 := 0; i2 < 3; i2++ {
		for j2 := 0; j2 < 3; j2++ {
			if board[areaX*3+i2][areaY*3+j2] == byte('0'+num) {
				return false
			}
		}
	}

	return true
}
