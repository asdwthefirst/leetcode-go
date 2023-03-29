/**
  @author: jiangxi
  @since: 2022/12/28
  @desc: //TODO
**/
package greed

func partitionLabels(s string) []int {
	str := []byte(s)
	lastPos := make(map[byte]int)
	for i := len(str) - 1; i >= 0; i-- {
		if _, ok := lastPos[str[i]]; !ok {
			lastPos[str[i]] = i
		}
	}
	left, right := 0, 0
	ans := []int{}
	for i := 0; i < len(str); i++ {
		for i <= right && i < len(str) {
			if lastPos[str[i]] > right {
				right = lastPos[str[i]]
			}
			i++
		}
		ans = append(ans, right-left+1)
		if i < len(str) {
			left, right = i, i
			i--
		}
	}
	return ans
}
