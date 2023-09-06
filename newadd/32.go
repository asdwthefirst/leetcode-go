/*
*

	@author: jiangxi
	@since: 2023/8/28
	@desc: //TODO

*
*/
package newadd

func longestValidParentheses(s string) int {
	max := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		// curS:=s[i:i+1]
		if s[i:i+1] == "(" {
			dp[i] = 0
		} else if s[i-1:i] == "(" {
			dp[i] = 2
			// if i-2>=0{
			//     dp[i]+=dp[i-2]
			// }
		} else if s[i-1:i] == ")" {
			pos := i - 1 - dp[i-1]
			if pos >= 0 && s[pos:pos+1] == "(" {
				dp[i] = dp[i-1] + 2
			}
		}
		for i-dp[i] >= 0 && dp[i-dp[i]] != 0 {
			dp[i] += dp[i-dp[i]]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
