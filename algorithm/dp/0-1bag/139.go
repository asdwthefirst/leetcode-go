/**
  @author: jiangxi
  @since: 2023/1/4
  @desc: //TODO
**/
package __1bag

//求解顺序的完全背包问题时，对物品的迭代应该放在最里层，对背包的迭代放在外层，只有这样才能让物品按一定顺序放入背包中。
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for _, word := range wordDict {
			if i-len(word) >= 0 && s[i-len(word):i] == word {
				dp[i] = dp[i-len(word)] || dp[i]
			}
		}
	}
	return dp[n]
}
