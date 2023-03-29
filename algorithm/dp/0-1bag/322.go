/**
  @author: jiangxi
  @since: 2023/1/4
  @desc: //TODO
**/
package __1bag

import (
	"math"
	"sort"
)

//这是我自己写的，建议用完全背包问题：
//因为硬币可以重复使用，因此这是一个完全背包问题。完全背包只需要将 0-1 背包的逆序遍历 dp 数组改为正序遍历即可。
//（这样向上取的时候，就以及涵盖了使用了0个i，1个i，。。。的情况，不需要逆序。等第二次写再写出来
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt / 2 //因为加法操作溢出了。。。/2就不会了
	}
	sort.Ints(coins)
	dp[0] = 0
	for i := len(coins) - 1; i >= 0; i-- {
		for j := amount; j >= coins[i]; j-- {
			for z := 0; z*coins[i] <= j; z++ {
				//fmt.Println(coins[i], z)
				dp[j] = min322(dp[j-z*coins[i]]+z, dp[j])
				//fmt.Println(j, ":", dp[j])
			}
		}
		//fmt.Println(coins[i], dp)
	}
	if dp[amount] >= math.MaxInt/2 {
		return -1
	}
	return dp[amount]

}

//dp[j]表示前i种钱币，组成j元钱的最小数量，
//coin,

func min322(a, b int) int {
	if a > b {
		return b
	}
	return a
}
