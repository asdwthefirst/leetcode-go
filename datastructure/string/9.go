/**
  @author: jiangxi
  @since: 2022/11/12
  @desc: //TODO
**/
package string

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xStr := fmt.Sprint(x)
	if len(xStr) == 1 {
		return true
	}
	for i, j := 0, len(xStr)-1; i <= j; i, j = i-1, j+1 {
		if xStr[i] != xStr[j] {
			return false
		}
	}
	return true
}
