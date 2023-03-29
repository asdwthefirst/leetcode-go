/**
  @author: jiangxi
  @since: 2023/1/4
  @desc: //TODO
**/
package stocktrade

//dp[i][0]表示今天结束时持有，dp[i][1]表示今天结束时不持有
func maxProfit714(prices []int, fee int) int {
	n := len(prices)
	//dp := make([][2]int, n)
	//dp[0][0] = -prices[0]
	//dp[0][1] = 0
	dp0 := -prices[0]
	dp1 := 0
	//max := 0
	for i := 1; i < n; i++ {
		dp0, dp1 = max714(dp1-prices[i], dp0), max714(dp1, dp0+prices[i]-fee)
		//if dp1 > max {
		//	max = dp1
		//}
	}
	return max714(dp0, dp1)

}

func max714(a, b int) int {
	if a > b {
		return a
	}
	return b
}
