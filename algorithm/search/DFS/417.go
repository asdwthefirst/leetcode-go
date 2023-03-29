/**
  @author: jiangxi
  @since: 2022/12/29
  @desc: //TODO
**/
package DFS

//太复杂了，应该分别判断可流alt和pac，然后
//var visited417 [][]bool
//var isResult [][]bool
//var xLen417, yLen417 int
//var reachPac, reachAtl bool
//
//func pacificAtlantic(heights [][]int) (ret [][]int) {
//	xLen417, yLen417 = len(heights), len(heights[0])
//	//visited417 = make([][]bool, xLen417)
//	//for i := range visited417 {
//	//	visited417[i] = make([]bool, yLen417)
//	//}
//	isResult = make([][]bool, xLen417)
//	for i := range isResult {
//		isResult[i] = make([]bool, yLen417)
//	}
//	isResult[0][yLen417-1], isResult[xLen417-1][0] = true, true
//	for i := 0; i < xLen417; i++ {
//		for j := 0; j < yLen417; j++ {
//			if isResult[i][j] == false {
//				visited417 = make([][]bool, xLen417)
//				for i := range visited417 {
//					visited417[i] = make([]bool, yLen417)
//				}
//				reachPac, reachAtl = false, false
//				canAdd := dfs417(heights, i, j)
//				if canAdd || reachPac && reachAtl {
//					isResult[i][j] = true
//				}
//			}
//		}
//	}
//
//	for i := 0; i < xLen417; i++ {
//		for j := 0; j < yLen417; j++ {
//			if isResult[i][j] {
//				ret = append(ret, []int{i, j})
//			}
//
//		}
//
//	}
//	return
//
//}
//
//func dfs417(heights [][]int, x, y int) bool {
//	visited417[x][y] = true
//	posNode := [][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
//	if x == 0 || y == 0 {
//		reachPac = true
//	}
//	if x == xLen417-1 || y == yLen417-1 {
//		reachAtl = true
//	}
//	for _, node := range posNode {
//		inX, inY := node[0], node[1]
//		if inX >= 0 && inY >= 0 && inX < xLen417 && inY < yLen417 && heights[inX][inY] <= heights[x][y] {
//			goBack := false
//			if visited417[inX][inY] == false {
//				goBack = dfs417(heights, inX, inY)
//			}
//			if goBack {
//				return goBack
//			}
//			if isResult[inX][inY] {
//				isResult[x][y] = true
//				return true
//			}
//		}
//	}
//	return false
//}

var visitedPac [][]bool
var visitedAtl [][]bool

//var visited417 [][]int
var reachPac [][]bool
var reachAtl [][]bool
var xLen417, yLen417 int
var Pac, Atl = 0, 1

//还是没有办法，用一个visited表示两个状态。这个不知道为什么有错误的，改天再看吧

func pacificAtlantic(heights [][]int) (ret [][]int) {
	xLen417, yLen417 = len(heights), len(heights[0])
	visitedPac = make([][]bool, xLen417)
	for i := range visitedPac {
		visitedPac[i] = make([]bool, yLen417)
	}
	visitedAtl = make([][]bool, xLen417)
	for i := range visitedAtl {
		visitedAtl[i] = make([]bool, yLen417)
	}
	reachPac = make([][]bool, xLen417)
	for i := range reachPac {
		reachPac[i] = make([]bool, yLen417)
	}
	reachAtl = make([][]bool, xLen417)
	for i := range reachAtl {
		reachAtl[i] = make([]bool, yLen417)
	}
	for i := 0; i < xLen417; i++ {
		reachAtl[i][0] = true
		visitedAtl[i][0] = true
		reachPac[i][yLen417-1] = true
		visitedPac[i][yLen417-1] = true
	}
	for i := 0; i < yLen417; i++ {
		reachAtl[0][i] = true
		reachPac[xLen417-1][i] = true
		visitedAtl[0][i] = true
		visitedPac[xLen417-1][i] = true
	}
	for i := 0; i < xLen417; i++ {
		for j := 0; j < yLen417; j++ {
			if !visitedAtl[i][j] {
				//fmt.Println("head:", i, j)
				dfsAtl(heights, i, j)
			}
			//dfs417(heights, i, j)
		}
	}
	for i := 0; i < xLen417; i++ {
		for j := 0; j < yLen417; j++ {
			if !visitedPac[i][j] {
				//fmt.Println("head:", i, j)
				dfsPac(heights, i, j)
			}
			//dfs417(heights, i, j)
		}
	}

	for i := 0; i < xLen417; i++ {
		for j := 0; j < yLen417; j++ {
			if reachPac[i][j] && reachAtl[i][j] {
				ret = append(ret, []int{i, j})
			}

		}

	}
	return

}

func dfsAtl(heights [][]int, x, y int) {
	visitedAtl[x][y] = true
	posNode := [][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}

	for _, node := range posNode {
		inX, inY := node[0], node[1]
		if inX >= 0 && inY >= 0 && inX < xLen417 && inY < yLen417 && heights[inX][inY] <= heights[x][y] {
			if !visitedAtl[inX][inY] {
				dfsAtl(heights, inX, inY)
			}
			if reachAtl[inX][inY] == true {
				reachAtl[x][y] = true
				return
			}
		}
	}
	return
}
func dfsPac(heights [][]int, x, y int) {
	visitedPac[x][y] = true
	posNode := [][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}

	for _, node := range posNode {
		inX, inY := node[0], node[1]
		if inX >= 0 && inY >= 0 && inX < xLen417 && inY < yLen417 && heights[inX][inY] <= heights[x][y] {
			if !visitedPac[inX][inY] {
				dfsPac(heights, inX, inY)
			}
			if reachPac[inX][inY] == true {
				reachPac[x][y] = true
				return
			}
		}
	}
	return
}
