/**
  @author: jiangxi
  @since: 2023/1/4
  @desc: //TODO
**/
package longestincreasingsubsequence

//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
//子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
//
//
//示例 1：
//
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

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
