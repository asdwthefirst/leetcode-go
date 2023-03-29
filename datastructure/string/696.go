/**
  @author: jiangxi
  @since: 2022/11/12
  @desc: //TODO
**/
package string

func countBinarySubstrings(s string) int {
	curLen := 0
	res := 0
	var counts []int
	for i := range s {
		if i == 0 || s[i] == s[i-1] {
			curLen++
		} else {
			counts = append(counts, curLen)
			curLen = 1
		}
	}
	counts = append(counts, curLen)
	for i := 0; i < len(counts)-1; i++ {
		res += min(counts[i], counts[i+1])
	}
	return res

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
