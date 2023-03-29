/**
  @author: jiangxi
  @since: 2023/1/4
  @desc: //TODO
**/
package longestincreasingsubsequence

import (
	"fmt"
	"sort"
)

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] < pairs[j][1] { //好像不管是根据结尾排序还是根据开头排序都可行
			return true
		}
		return false
	})
	fmt.Println(pairs)
	n := len(pairs)
	dp := make([]int, n)
	max := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if pairs[j][1] < pairs[i][0] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
