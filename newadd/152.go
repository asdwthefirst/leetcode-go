/*
*

	@author: jiangxi
	@since: 2023/8/28
	@desc: //TODO

*
*/
package newadd

// 分正负讨论，维护一个dpmin和dpmax,0是min,1是max
func maxProduct(nums []int) int {
	// dp:=make([][2]int,len(nums))
	dp0, dp1 := nums[0], nums[0]
	// dp[0][0],dp[0][1]=nums[i],nums[i]
	max := dp1
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			dp0, dp1 = 0, 0
		}
		if nums[i] > 0 {
			dp0, dp1 = min152(dp0*nums[i], nums[i]), max152(dp1*nums[i], nums[i])
		} else {
			dp0, dp1 = min152(dp1*nums[i], nums[i]), max152(dp0*nums[i], nums[i])
		}
		// fmt.Println(i,dp0,dp1)
		if dp1 > max {
			max = dp1
		}
	}
	return max

}

func max152(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min152(a, b int) int {
	if a < b {
		return a
	}
	return b
}
