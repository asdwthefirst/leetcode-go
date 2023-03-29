/**
  @author: jiangxi
  @since: 2022/12/28
  @desc: //TODO
**/
package binary

import (
	"fmt"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[len(letters)-1] < target {
		return letters[0]
	}
	left, right := 0, len(letters)-1
	for left < right {
		fmt.Println(left, right)
		mid := left + (right-left)/2 //找右边界，特殊处理一下+1
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return letters[left]
}
