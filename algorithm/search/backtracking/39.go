/**
  @author: jiangxi
  @since: 2022/12/30
  @desc: //TODO
**/
package backtracking

//回头再看看，处理有点蒙的感觉
func combinationSum(candidates []int, target int) (ret [][]int) {
	visited := make([]bool, len(candidates))
	for i := 0; i < len(candidates); i++ {
		if candidates[i] <= target {
			backtrack39(candidates, i, visited, target, &ret, []int{})
		}
	}
	return

}

func backtrack39(candidates []int, index int, visited []bool, target int, ret *[][]int, choosed []int) {
	//visited[index] = true
	choosed = append(choosed, candidates[index])
	target -= candidates[index]
	if target == 0 {
		*ret = append(*ret, append([]int{}, choosed...))
	} else {
		for i := index; i < len(candidates); i++ { //没想得特别清楚，，但是通过了，回头再看看把
			if candidates[i] <= target {
				backtrack39(candidates, i, visited, target, ret, choosed)
			}
		}
	}
	//visited[index] = false
	return
}
