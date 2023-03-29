/**
  @author: jiangxi
  @since: 2022/11/11
  @desc: //TODO
**/
package string

func longestPalindrome(s string) int {
	charMap := make(map[rune]int)
	for _, ch := range s {
		charMap[ch]++
	}
	res := 0
	for _, v := range charMap {
		if v%2 == 0 {
			res += v
		} else {
			res += v - 1
		}
	}
	if res < len(s) {
		res++
	}
	return res
}
