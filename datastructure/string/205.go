/**
  @author: jiangxi
  @since: 2022/11/12
  @desc: //TODO
**/
package string

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap := make(map[uint8]uint8)
	reverseCharMap := make(map[uint8]uint8)

	for i := 0; i < len(s); i++ {
		fmt.Println(i, charMap, s[i], t[i])
		sCh := s[i]
		tCh := t[i]
		if tChEx, ok := charMap[sCh]; ok {
			if tChEx != tCh {
				return false
			}
		}
		if sChEx, ok := reverseCharMap[tCh]; ok {
			if sChEx != sCh {
				return false
			}
		}
		charMap[sCh] = tCh
		reverseCharMap[tCh] = sCh
	}
	return true

}
