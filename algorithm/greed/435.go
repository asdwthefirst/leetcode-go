/**
  @author: jiangxi
  @since: 2022/12/27
  @desc: //TODO
**/
package greed

import "sort"

//每次在可选区间选结尾最靠前的区间，毫无疑问留给后面的空间会更大，是我自己想不到的
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans, right := 1, intervals[0][1]
	for _, in := range intervals[1:] {
		if in[0] >= right {
			ans++
			right = in[1]
		}
	}
	return len(intervals) - ans
}
