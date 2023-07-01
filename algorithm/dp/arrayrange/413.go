/*
*

	@author: jiangxi
	@since: 2023/1/3
	@desc: //TODO

*
*/
package arrayrange

//

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	//【1，2，3，4，5，6，7，8】
	//dp[i][j]是i，j是否是等差。
	//i-1,i,i+1,等差，则dp[a][i+1]=dp[a][i]
	//dp2[i+1]为以i+1为结尾的等差数组数量，i-1,i,i+1,等差,则dp2[i+1]=dp2[i]+1,否则dp2[i+1]=0
	//用一个sum统计。只需要存上一位的值。
	dp := 0
	sum := 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp = dp + 1
		} else {
			dp = 0
		}
		sum += dp
	}
	return sum
}
