/**
  @author: jiangxi
  @since: 2023/1/4
  @desc: //TODO
**/
package cutint

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if canDecode(s[i : i+1]) {
			if i-1 >= 0 {
				dp[i] += dp[i-1]
			} else {
				dp[i] += 1
			}
		}
		if i-1 >= 0 && canDecode(s[i-1:i+1]) {
			if i-2 >= 0 {
				dp[i] += dp[i-2]
			} else {
				dp[i] += 1
			}
		}
	}
	fmt.Println(dp)
	return dp[n-1]

}

func canDecode(s string) bool {
	if len(s) == 0 || len(s) > 2 {
		return false
	}
	if s[:1] == "0" {
		return false
	}
	num, _ := strconv.ParseInt(s, 10, 64)
	if num <= 26 && num >= 0 {
		return true
	}
	return false
}
