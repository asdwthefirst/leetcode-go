/**
  @author: jiangxi
  @since: 2022/12/30
  @desc: //TODO
**/
package backtracking

func permute(nums []int) (ret [][]int) {
	for i := 0; i < len(nums); i++ {
		visited := make([]bool, len(nums))
		backtrack46(nums, visited, i, []int{}, &ret)
	}
	return

}

func backtrack46(nums []int, visited []bool, index int, permuted []int, ret *[][]int) {
	permuted = append(permuted, nums[index])
	visited[index] = true
	if len(permuted) == len(nums) {
		*ret = append(*ret, permuted)
	} else {
		for i := 0; i < len(nums); i++ {
			if !visited[i] {
				backtrack46(nums, visited, i, permuted, ret)
			}
		}
	}
	visited[index] = false
	return
}
