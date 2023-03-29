/**
  @author: jiangxi
  @since: 2022/12/19
  @desc: //TODO
**/
package doublepointer

func validPalindrome(s string) bool {
	for pa, pb := 0, len(s)-1; pa <= pb; pa, pb = pa+1, pb-1 {
		if s[pa] != s[pb] {
			return palindrome680(s, pa+1, pb) || palindrome680(s, pa, pb-1)
		}
	}
	return true
}

func palindrome680(s string, pa, pb int) bool {
	for pa <= pb {
		if s[pa] != s[pb] {
			return false
		}
		pa++
		pb--
	}
	return true
}
