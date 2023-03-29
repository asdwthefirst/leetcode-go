/**
  @author: jiangxi
  @since: 2023/1/4
  @desc: //TODO
**/
package __1bag

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	W := (sum + target) / 2
	if W*2 != sum+target || W < 0 {
		return 0
	}
	dp := make([]int, W+1)
	n := len(nums)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := W; j >= nums[i]; j-- {
			dp[j] = dp[j] + dp[j-nums[i]] //倒序使得dp[j-nums[i]]给出的是考虑i-1个数的情况
		}
	}
	return dp[W]
}

//target=target,i，dp[i][j]表示前i个数，sum为j的个数，j的范围是-sum--+sum，最终结果是dp[target]
//dp[i][j]=dp[i-1][j-nums[i]]+dp[i-1][j+nums[i]]
//以上是我的想法，但实际上这个问题实际上是选中n个数使得choosedSum-（sum-choosedSum）=target，choosedSum=(sum+target)/2的背包问题
//dp[j]表示前i个数，sum为j的个数，j的范围是choosedsum--0，dp[j]=dp[j-nums[i]]+dp[j]
