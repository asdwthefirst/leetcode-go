/**
  @author: jiangxi
  @since: 2023/1/5
  @desc: //TODO
**/
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
