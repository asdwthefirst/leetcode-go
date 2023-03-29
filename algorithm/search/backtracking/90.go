/**
  @author: jiangxi
  @since: 2022/12/30
  @desc: //TODO
**/
package backtracking

import "sort"

func subsetsWithDup(nums []int) (ret [][]int) {
	sort.Ints(nums)
	visited := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			backtrack90(nums, i, visited, &ret, []int{})
		}
	}
	ret = append(ret, []int{})
	return

}

func backtrack90(nums []int, index int, visited []bool, ret *[][]int, choosed []int) {
	visited[index] = true
	choosed = append(choosed, nums[index])
	*ret = append(*ret, append([]int{}, choosed...))
	for i := index + 1; i < len(nums); i++ {
		if !visited[i] && (nums[i] != nums[i-1] || visited[i-1]) {
			backtrack90(nums, i, visited, ret, choosed)
		}
	}
	visited[index] = false
	return
}
