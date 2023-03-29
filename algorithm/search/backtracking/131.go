/**
  @author: jiangxi
  @since: 2022/12/30
  @desc: //TODO
**/
package backtracking

func partition(s string) (ret [][]string) {
	backtrack131(s, &ret, []string{})
	return

}

func backtrack131(s string, ret *[][]string, pals []string) {
	//fmt.Println(s, pals)
	if len(s) == 0 {
		*ret = append(*ret, append([]string{}, pals...))
	} else {
		for i := 1; i <= len(s); i++ {
			if isPalindrome(s[:i]) {
				//fmt.Println(s[:i], s[i:])
				pals = append(pals, s[:i])
				backtrack131(s[i:], ret, pals)
				pals = pals[:len(pals)-1]
			}
		}
	}
	return
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
