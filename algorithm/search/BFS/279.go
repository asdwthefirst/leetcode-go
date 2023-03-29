/**
  @author: jiangxi
  @since: 2022/12/29
  @desc: //TODO
**/
package BFS

func numSquares(n int) int {
	squares := generateSquares(n)
	pathLen := make([]int, n+1)
	pathLen[n] = 0
	var queue []int
	queue = append(queue, n)
	for len(queue) != 0 {
		curNum := queue[0]
		queue = queue[1:]
		if curNum == 0 {
			return pathLen[curNum]
		}
		for i := len(squares) - 1; i >= 0; i-- {
			if squares[i] <= curNum {
				newNum := curNum - squares[i]
				if pathLen[newNum] == 0 {
					pathLen[curNum-squares[i]] = pathLen[curNum] + 1
					queue = append(queue, newNum)
				}
			}
		}
	}
	return 0
}

//神奇的算法～
func generateSquares(n int) (ans []int) {
	square := 1
	diff := 3
	for square <= n {
		ans = append(ans, square)
		square += diff
		diff += 2
	}
	return
}
