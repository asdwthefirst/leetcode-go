/**
  @author: jiangxi
  @since: 2022/12/29
  @desc: //TODO
**/
package BFS

//在二维矩阵中搜索，什么时候用BFS，什么时候用DFS？
//1.如果只是要找到某一个结果是否存在，那么DFS会更高效。因为DFS会首先把一种可能的情况尝试到底，
//才会回溯去尝试下一种情况，只要找到一种情况，就可以返回了。但是BFS必须所有可能的情况同时尝试，
//在找到一种满足条件的结果的同时，也尝试了很多不必要的路径；
//2.如果是要找所有可能结果中最短的，那么BFS回更高效。因为DFS是一种一种的尝试，
//在把所有可能情况尝试完之前，无法确定哪个是最短，所以DFS必须把所有情况都找一遍，
//才能确定最终答案（DFS的优化就是剪枝，不剪枝很容易超时）。而BFS从一开始就是尝试所有情况，
//所以只要找到第一个达到的那个点，那就是最短的路径，可以直接返回了，其他情况都可以省略了，
//所以这种情况下，BFS更高效。
//

func shortestPathBinaryMatrix(grid [][]int) int {
	if len(grid) < 1 {
		return -1
	}
	n := len(grid)
	if grid[0][0] != 0 || grid[n-1][n-1] != 0 {
		return -1
	}
	if n == 1 {
		return 1
	}
	pathLen := make([][]int, n)
	for i := range pathLen {
		pathLen[i] = make([]int, n)
	}
	pathLen[0][0] = 1
	var nodeLst [][2]int
	nodeLst = append(nodeLst, [2]int{0, 0})
	for len(nodeLst) >= 0 {
		curNode := nodeLst[0] //queue pop
		nodeLst = nodeLst[1:]
		x, y := curNode[0], curNode[1]
		if x == n-1 && y == n-1 {
			return pathLen[x][y]
		}
		posNode := [][2]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x - 1, y + 1}, {x + 1, y + 1}, {x, y - 1}, {x - 1, y - 1}, {x + 1, y - 1}}
		for _, node := range posNode {
			inX := node[0]
			inY := node[1]
			if inX < n && inX >= 0 && inY < n && inY >= 0 && grid[inX][inY] == 0 && pathLen[inX][inY] == 0 {
				pathLen[inX][inY] = pathLen[x][y] + 1
				nodeLst = append(nodeLst, [2]int{inX, inY})
			}
		}
	}
	return -1

}
