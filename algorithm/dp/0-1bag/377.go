/*
*

	@author: jiangxi
	@since: 2023/1/4
	@desc: //TODO

*
*/
package __1bag

//给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
//
//题目数据保证答案符合 32 位整数范围。
//请注意，顺序不同的序列被视作不同的组合。

func combinationSum4(nums []int, target int) int {
	//好神奇，这种不能从数组的从前往后考虑，就把数组放在内层
	dp := make([]int, target+1)
	dp[0] = 0
	for i := 0; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[num] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func combinationSum40701(nums []int, target int) int {
	//那这不是组合问题。。。排列问题
	//dp[i]为组成i可能的组合数
	//for j in range coin, dp[i]+=dp[i-coin[j]]
	if target == 0 {
		return 1
	}
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i-nums[j] >= 0 {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]

}
