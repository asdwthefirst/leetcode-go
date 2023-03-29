/**
  @author: jiangxi
  @since: 2023/1/4
  @desc: //TODO
**/
package __1bag

func combinationSum4(nums []int, target int) int {
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
