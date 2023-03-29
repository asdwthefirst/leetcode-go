/**
  @author: jiangxi
  @since: 2022/12/15
  @desc: //TODO
**/
package arraymatrix

import "fmt"

func findShortestSubArray(nums []int) int {
	start := make(map[int]int)
	count := make(map[int]int)
	for i, v := range nums {
		count[v]++
		if _, ok := start[v]; !ok {
			start[v] = i
		}
	}
	max := 0
	ansInt := []int{}
	for i, v := range count {
		if v == max {
			ansInt = append(ansInt, i)
		}
		if v > max {
			ansInt = ansInt[:0]
			max = v
			ansInt = append(ansInt, i)
		}
	}
	fmt.Println(ansInt)
	ans := []int{}
	for _, ansInt := range ansInt {
		for i := len(nums) - 1; i >= 0; i-- {
			if nums[i] == ansInt {
				ans = append(ans, i-start[ansInt]+1)
				break
			}
		}
	}
	min := ans[0]
	for _, v := range ans {
		if v < min {
			min = v
		}
	}
	return min

}
