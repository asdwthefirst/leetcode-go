/**
  @author: jiangxi
  @since: 2023/1/4
  @desc: //TODO
**/
package stocktrade

import "fmt"

//分成三种状态算dp
//dp[i][0],持有，dp[i][1]，不持有，今天卖了，dp[i][2]，不持有，不是今天卖的

func maxProfit309(prices []int) int {
	n := len(prices)

	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = max309(dp[i-1][0], dp[i-1][2]-prices[i]) //1。昨天持有，今天继续持有，2。昨天不持有不是昨天卖的，今天买了
		dp[i][1] = dp[i-1][0] + prices[i]                   //昨天持有，今天卖了
		dp[i][2] = max309(dp[i-1][1], dp[i-1][2])           //昨天也不持有
	}
	fmt.Println(dp)
	return max309(max309(dp[n-1][0], dp[n-1][1]), dp[n-1][2])

}

func max309(a, b int) int {
	if a > b {
		return a
	}
	return b
}
