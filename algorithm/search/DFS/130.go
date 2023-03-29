/**
  @author: jiangxi
  @since: 2022/12/29
  @desc: //TODO
**/
package DFS

var visited130 [][]bool
var curNodeLst [][2]int
var notChange bool
var xLen130, yLen130 int

func solve(board [][]byte) {
	xLen130, yLen130 = len(board), len(board[0])
	visited130 = make([][]bool, xLen130)
	for i := range visited130 {
		visited130[i] = make([]bool, yLen130)
	}
	for i := 0; i < xLen130; i++ {
		for j := 0; j < yLen130; j++ {
			if board[i][j] == 'O' && visited130[i][j] == false {
				curNodeLst = curNodeLst[:0]
				notChange = false
				dfs130(board, i, j)
				// fmt.Println(notChange,curNodeLst)
				if notChange == false {
					for _, node := range curNodeLst {
						board[node[0]][node[1]] = 'X'
					}
				}
			}
		}
	}
	return

}

func dfs130(board [][]byte, x, y int) {
	// fmt.Println(x,y)
	visited130[x][y] = true
	curNodeLst = append(curNodeLst, [2]int{x, y})
	if x == 0 || y == 0 || x == xLen130-1 || y == yLen130-1 {
		notChange = true
	}
	posNode := [][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
	for _, node := range posNode {
		inX, inY := node[0], node[1]
		if inX >= 0 && inY >= 0 && inX < xLen130 && inY < yLen130 && board[inX][inY] == 'O' && visited130[inX][inY] == false {
			dfs130(board, inX, inY)
		}
	}
	return
}
