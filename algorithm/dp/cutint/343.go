/**
  @author: jiangxi
  @since: 2023/1/3
  @desc: //TODO
**/
package cutint

import "fmt"

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		max := 0
		for j := 1; j < i; j++ {
			//fmt.Println(i, j, i-j, dp[i-j])
			pos1, pos2 := dp[i-j]*j, (i-j)*j
			if pos1 > max {
				max = pos1
			}
			if pos2 > max {
				max = pos2
			}
		}
		dp[i] = max
		fmt.Println(dp)
	}
	return dp[n]
}
