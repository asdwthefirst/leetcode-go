/**
  @author: jiangxi
  @since: 2022/12/29
  @desc: //TODO
**/
package DFS

var visited [][]bool
var xLen, yLen int
var area int

func maxAreaOfIsland(grid [][]int) int {
	xLen, yLen = len(grid), len(grid[0])
	maxArea := 0
	visited = make([][]bool, xLen)
	for i := range visited {
		visited[i] = make([]bool, yLen)
	}
	for i := 0; i < xLen; i++ {
		for j := 0; j < yLen; j++ {
			if visited[i][j] == false && grid[i][j] == 1 {
				area = 0
				dfs695(grid, i, j)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea

}

func dfs695(grid [][]int, x, y int) {
	visited[x][y] = true
	area++
	posNode := [][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
	for _, node := range posNode {
		inX, inY := node[0], node[1]
		if inX >= 0 && inY >= 0 && inX < xLen && inY < yLen && grid[inX][inY] == 1 && visited[inX][inY] == false {
			dfs695(grid, inX, inY)
		}
	}
	return
}
