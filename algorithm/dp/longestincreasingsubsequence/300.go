/**
  @author: jiangxi
  @since: 2023/1/4
  @desc: //TODO
**/
package longestincreasingsubsequence

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	max := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
