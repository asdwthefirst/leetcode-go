/**
  @author: jiangxi
  @since: 2023/1/3
  @desc: //TODO
**/
package cutint

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	squares := generateSquares(n)
	for i := 1; i < n+1; i++ {
		min := math.MaxInt
		for _, square := range squares {
			if square <= i && dp[i-square]+1 < min {
				dp[i] = dp[i-square] + 1
			}
		}
		dp[i] = min
		fmt.Println(dp)
	}
	return dp[n]

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
