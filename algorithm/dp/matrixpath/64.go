/**
  @author: jiangxi
  @since: 2023/1/3
  @desc: //TODO
**/
package matrixpath

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	minPath := make([][]int, m)
	for i := range minPath {
		minPath[i] = make([]int, n)
	}
	minPath[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		minPath[0][i] = minPath[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		minPath[i][0] = minPath[i-1][0] + grid[i][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			minPath[i][j] = grid[i][j] + min64(minPath[i-1][j], minPath[i][j-1])
		}
	}
	return minPath[m-1][n-1]

}

func min64(a, b int) int {
	if a > b {
		return b
	}
	return a
}
