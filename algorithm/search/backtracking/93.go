/**
  @author: jiangxi
  @since: 2022/12/30
  @desc: //TODO
**/
package backtracking

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	var ret []string
	backtrack93(&ret, s, 0, "", 0)
	return ret
}

func backtrack93(ret *[]string, s string, index int, prefix string, hasPos int) {
	fmt.Println(prefix)
	if index == len(s) && hasPos == 4 {
		*ret = append(*ret, prefix[:len(prefix)-1])
		return
	}
	if index == len(s) || hasPos == 4 {
		return
	}
	if s[index:index+1] == "0" {
		backtrack93(ret, s, index+1, prefix+"0.", hasPos+1)
		return
	}
	for i := 1; i <= 3; i++ {
		if isValidIpNum(s, index, index+i) {
			backtrack93(ret, s, index+i, prefix+s[index:index+i]+".", hasPos+1)
		}
	}
	return

}

func isValidIpNum(s string, start, end int) bool {
	if end > len(s) {
		return false
	}
	numStr := s[start:end]
	if len(numStr) > 3 || len(numStr) == 0 {
		return false
	}
	num, _ := strconv.ParseInt(numStr, 10, 64)
	if num > 255 {
		return false
	}
	return true
}
