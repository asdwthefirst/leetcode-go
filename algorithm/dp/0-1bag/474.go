/**
  @author: jiangxi
  @since: 2023/1/4
  @desc: //TODO
**/
package __1bag

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 0
	for _, str := range strs {
		zeros, ones := countZeroOne(str)
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = max474(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}
	return dp[m][n]

}

func countZeroOne(str string) (m, n int) {
	for _, ch := range str {
		if ch == '0' {
			m++
		}
		if ch == '1' {
			n++
		}
	}
	return
}

func max474(a, b int) int {
	if a > b {
		return a
	}
	return b
}