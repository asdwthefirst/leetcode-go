/**
  @author: jiangxi
  @since: 2023/1/5
  @desc: //TODO
**/
package stocktrade

import "fmt"

//困难一把过，我尊牛逼

//分类：资产
//0.结束时候持有，已卖出一次（算卖出值，与最大值比较,,,dp[i][0]=max123(dp[i-1][0],dp[i-1][2]-prices[i])
//1.结束时持有，未卖出过,,,dp[i][1]=max123(dp[i-1][1],-prices[i])
//2.结束时不持有，卖出一次,,,dp[i][2]=max123(dp[i-1][2],dp[i-1][1]+prices[i])
//3.结束时不持有，卖出两次,,,dp[i][3]=max123(dp[i-1][3],dp[i-1][0]+prices[i])
////3.结束时不持有，未买入过，（那一定是0）
func maxProfit123(prices []int) int {
	n := len(prices)
	dp0 := -prices[0]
	dp1 := -prices[0]
	dp2 := 0
	dp3 := 0
	for i := 1; i < n; i++ {
		dp0, dp3 = max123(dp0, dp2-prices[i]), max123(dp3, dp0+prices[i])
		dp2 = max123(dp2, dp1+prices[i])
		dp1 = max123(dp1, -prices[i])
		fmt.Println(i, ":", dp0, dp1, dp2, dp3)
	}
	return max123(max123(max123(dp0, dp1), dp2), dp3)

}

func max123(a, b int) int {
	if a > b {
		return a
	}
	return b
}
