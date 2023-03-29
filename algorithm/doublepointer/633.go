/**
  @author: jiangxi
  @since: 2022/12/19
  @desc: //TODO
**/
package doublepointer

import "math"

//需剪枝
func judgeSquareSum(c int) bool {
	for a := 0; a*a <= c; a++ {
		rt := math.Sqrt(float64(c - a*a))
		if rt == math.Floor(rt) {
			return true
		}
	}
	return false

}
