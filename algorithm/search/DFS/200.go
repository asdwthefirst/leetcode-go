/**
  @author: jiangxi
  @since: 2022/12/29
  @desc: //TODO
**/
package DFS

import "fmt"

var visited200 [][]bool
var xLen200, yLen200 int

func numIslands(grid [][]byte) int {
	xLen200, yLen200 = len(grid), len(grid[0])
	visited200 = make([][]bool, xLen200)
	for i := range visited200 {
		visited200[i] = make([]bool, yLen200)
	}
	ans := 0
	for i := 0; i < xLen200; i++ {
		for j := 0; j < yLen200; j++ {
			fmt.Println(i, j)
			if visited200[i][j] == false && grid[i][j] == '1' {
				fmt.Println("dfs:", i, j)
				dfs200(grid, i, j)
				ans++
			}
		}
	}
	return ans

}

func dfs200(grid [][]byte, x, y int) {
	visited200[x][y] = true
	posNode := [][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
	for _, node := range posNode {
		inX, inY := node[0], node[1]
		if inX >= 0 && inY >= 0 && inX < xLen200 && inY < yLen200 && grid[inX][inY] == '1' && visited200[inX][inY] == false {
			dfs200(grid, inX, inY)
		}
	}
	return
}
