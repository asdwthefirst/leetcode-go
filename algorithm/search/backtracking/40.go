/**
  @author: jiangxi
  @since: 2022/12/30
  @desc: //TODO
**/
package backtracking

import "sort"

func combinationSum2(candidates []int, target int) (ret [][]int) {
	sort.Ints(candidates)
	visited := make([]bool, len(candidates))
	for i := 0; i < len(candidates); i++ {
		if candidates[i] <= target && (i == 0 || candidates[i] != candidates[i-1]) {
			backtrack40(candidates, i, visited, target, &ret, []int{})
		}
	}
	return

}

func backtrack40(candidates []int, index int, visited []bool, target int, ret *[][]int, choosed []int) {
	visited[index] = true
	choosed = append(choosed, candidates[index])
	target -= candidates[index]
	if target == 0 {
		*ret = append(*ret, append([]int{}, choosed...))
	} else {
		for i := index + 1; i < len(candidates); i++ {
			if candidates[i] <= target && !visited[i] && (candidates[i] != candidates[i-1] || visited[i-1]) {
				backtrack40(candidates, i, visited, target, ret, choosed)
			}
		}
	}
	visited[index] = false
	return
}
