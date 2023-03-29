/**
  @author: jiangxi
  @since: 2023/1/3
  @desc: //TODO
**/
package arrayrange

//

func numberOfArithmeticSlices(nums []int) int {

	n := len(nums)
	if n < 3 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 0
	dp[1] = 0
	sum := 0
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
		}
		sum += dp[i]
	}
	return sum

}
