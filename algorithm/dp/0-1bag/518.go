/**
  @author: jiangxi
  @since: 2023/1/4
  @desc: //TODO
**/
package __1bag

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = dp[j-coin] + dp[j]
		}
	}
	return dp[amount]
}
