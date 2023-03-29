/**
  @author: jiangxi
  @since: 2023/1/4
  @desc: //TODO
**/
package __1bag

func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	W := sum / 2
	if W*2 != sum {
		return false
	}
	dp := make([]bool, W+1)
	n := len(nums)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := W; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]] //倒序使得dp[j-nums[i]]给出的是考虑i-1个数的情况
			if j == W && dp[j] == true {
				return true
			}
		}
	}
	return dp[W]
}
