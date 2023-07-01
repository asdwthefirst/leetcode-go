/*
*

	@author: jiangxi
	@since: 2023/1/5
	@desc: //TODO

*
*/
package stocktrade

//困难靠自己(别人的一点点启发）完成了，优秀

//分类：资产
//dp0[j],结束时候持有，已卖出j次，dp0[i][j]=max188(dp1[i-1][j-1]-prices[i],dp0[i-1][j]),,,dp0[i][0]=max188(dp[i-1][0],-prices[i])
//dp1[j],结束时不持有，已卖出j次,dp0[i][j]=max188(dp1[i-1][j],dp0[i-1][j]+prices[i]),,,dp1[i][0]=0

func maxProfit(k int, prices []int) int {
	n := len(prices)
	dp0 := make([]int, k+1)
	dp1 := make([]int, k+1)
	max := 0
	for j := 0; j < k+1; j++ {
		dp0[j] = -prices[0]
		dp1[j] = 0
	}
	for i := 1; i < n; i++ {
		dp0[0], dp1[0] = max188(dp0[0], -prices[i]), 0
		for j := 1; j <= k; j++ {
			dp0[j], dp1[j] = max188(dp1[j]-prices[i], dp0[j]), max188(dp1[j], dp0[j-1]+prices[i])
		}
		//fmt.Println(i, ":", dp0)
		//fmt.Println(i, ":", dp1)
	}
	for j := 0; j < k+1; j++ {
		if dp0[j] > max {
			max = dp0[j]
		}
		if dp1[j] > max {
			max = dp1[j]
		}
	}
	return max

}

func max188(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//dp[i][j]表示第i天，交易了j次(从买入开始算)，的最大money持有数，[0]持有[1]不持有
//i=0预留一位i-1,且dp[0][j][0]/[1]=0
//i=1,[0]=-prices[0]
//dp[i][j][0]=max（dp[i-1][j-1][1]-prices[i-1],dp[i-1][j][0])
//dp[i][j][1]=max(dp[i-1][j][0]+prices[i-1],dp[i-1][j][1])
//这样不对，这样没有购入的动力
//

func maxProfit0701(k int, prices []int) int {
	dp := make([][][2]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([][2]int, k+1)
	}
	for j := range dp[0] {
		dp[0][j][0] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max188(dp[i-1][j-1][1]-prices[i], dp[i-1][j][0])
			dp[i][j][1] = max188(dp[i-1][j][0]+prices[i], dp[i-1][j][1])
		}
	}
	max := 0
	for j := 0; j <= k; j++ {
		if dp[len(prices)-1][j][1] > max {
			max = dp[len(prices)-1][j][1]
		}
	}
	return max
}
