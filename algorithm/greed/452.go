/**
  @author: jiangxi
  @since: 2022/12/27
  @desc: //TODO
**/
package greed

import "sort"

//跟435就是一样的，我竟然看不出来。。。
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	ans, right := 1, points[0][1]
	for _, in := range points[1:] {
		if in[0] > right {
			ans++
			right = in[1]
		}
	}
	return ans

}
