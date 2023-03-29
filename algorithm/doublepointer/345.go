/**
  @author: jiangxi
  @since: 2022/12/19
  @desc: //TODO
**/
package doublepointer

import "strings"

func reverseVowels(s string) string {
	t := []byte(s) //string to char array
	i, j := 0, len(t)-1
	for i < j {
		for i < len(t) && !strings.Contains("aeiouAEIOU", string(t[i])) {
			i++
		}
		for j > 0 && !strings.Contains("aeiouAEIOU", string(t[j])) {
			j--
		}
		if i < j {
			t[i], t[j] = t[j], t[i]
			i++
			j--
		}

	}
	return string(t)

}
