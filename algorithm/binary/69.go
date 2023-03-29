/**
  @author: jiangxi
  @since: 2022/12/28
  @desc: //TODO
**/
package binary

import (
	"fmt"
	"math"
)

//溢出如何处理
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	left, right := 0, min69(x, math.MaxInt16)*2
	for left < right {
		fmt.Println(left, right)
		mid := left + (right-left)/2 + 1 //找右边界，特殊处理一下+1
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return right

}

func min69(a, b int) int {
	if a > b {
		return b
	}
	return a
}
