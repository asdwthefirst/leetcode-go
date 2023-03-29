/**
  @author: jiangxi
  @since: 2023/1/4
  @desc: //TODO
**/
package longestcommonsubsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1) //预留一位不用的0，避免转移的时候判断有没有越界。。。
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max1143(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func max1143(a, b int) int {
	if a > b {
		return a
	}
	return b
}
