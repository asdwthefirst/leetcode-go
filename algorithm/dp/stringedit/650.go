/*
*

	@author: jiangxi
	@since: 2023/1/5
	@desc: //TODO

*
*/
package stringedit

import (
	"math"
)

//减少枚举次数的细节值得思考:由于i必同时拥有因数j，i/j，则两者至少有一位小于i^(1/2)

// dp[i]表示完成i个a，至少进行的步数，dp[i]=min(dp[j]+i/j)，i/j整除
func minSteps(n int) int {
	dp := make([]int, n)
	dp[0] = 0 //dp[i]表示i+1个a
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt / 2
	}
	//优化前
	//for i := 1; i < n; i++ {
	//	for j := 0; j <= i/2; j++ {
	//		if (i+1)%(j+1) == 0 {
	//			dp[i] = min650(dp[j]+(i+1)/(j+1), dp[i])
	//		}
	//	}
	//}
	//优化后，时间缩短了6倍
	for i := 1; i < n; i++ {
		for j := 0; j < int(math.Sqrt(float64(i+1))); j++ {
			if (i+1)%(j+1) == 0 {
				dp[i] = min650(dp[j]+(i+1)/(j+1), dp[i])
				dp[i] = min650(dp[(i+1)/(j+1)-1]+j+1, dp[i])
			}
		}
	}
	return dp[n-1]
}

func min650(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// dp[1]=0，dp[n-1]为结果
// 初始状态，i>=1,dp[i]=i
// if i%j==0,dp[i]=min(dp[i],dp[j]+i/j)
func minSteps0701(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = i
	}
	for i := 2; i <= n; i++ {
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				dp[i] = min650(dp[i], dp[j]+i/j)
			}
		}
	}
	return dp[n]

}
