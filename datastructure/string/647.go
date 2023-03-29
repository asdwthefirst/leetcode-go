/**
  @author: jiangxi
  @since: 2022/11/12
  @desc: //TODO
**/
package string

func countSubstrings(s string) int {
	len := len(s)
	dp := make([][]bool, len)
	for i := range dp {
		dp[i] = make([]bool, len)
	}
	res := 0
	for i := 0; i < len; i++ {
		dp[i][i] = true
		res++
	}
	for i := len - 1; i >= 0; i-- {
		for j := i + 1; j < len; j++ {
			if s[i] == s[j] && (j == i+1 || dp[i+1][j-1]) {
				dp[i][j] = true
				res++
			} else {
				dp[i][j] = false
			}
		}
	}
	return res
}
