/*
*

	@author: jiangxi
	@since: 2022/12/30
	@desc: //TODO

*
*/
package backtracking

import "sort"

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

// 感觉思路没问题，但是出bug了，而且bug跟我46-47的剪枝操作毫无关系，不剪枝也出bug,明明append的时候是正常的，
// 为什么最后会出问题？找到原因了，在下面的注释里
//func PermuteUnique(nums []int) (ret [][]int) {
//	numMap := make(map[int]bool)
//	for i := 0; i < len(nums); i++ {
//		//for i := 0; i < 1; i++ {
//		_, ok := numMap[nums[i]]
//		if !ok {
//			visited := make([]bool, len(nums))
//			numMap[nums[i]] = true
//			backtrack47(nums, visited, i, []int{}, &ret)
//		}
//	}
//	return
//
//}
//
//func backtrack47(nums []int, visited []bool, index int, permuted []int, ret *[][]int) {
//	permuted = append(permuted, nums[index])
//	//permuted = append(permuted, index)
//
//	visited[index] = true
//	if len(permuted) == len(nums) {
//		*ret = append(*ret, append([]int{}, permuted...)) //非常关键，加入结果应该用新数组
//	} else {
//		numMap := make(map[int]bool)
//		for i := 0; i < len(nums); i++ {
//			//fmt.Println(permuted, numMap)
//			_, ok := numMap[nums[i]]
//			if !visited[i] && !ok {
//				numMap[nums[i]] = true
//				backtrack47(nums, visited, i, permuted, ret)
//			}
//		}
//	}
//	visited[index] = false
//	return
//}

func permuteUnique(nums []int) (ret [][]int) {
	sort.Ints(nums)
	visited := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		backtrack47(nums, visited, []int{}, i, &ret)
	}

}

func backtrack47(nums []int, visited []bool, currentRet []int, index int, ret *[][]int) {
	currentRet = append(currentRet, nums[index])
	visited[index] = true
	if len(currentRet) == len(nums) {
		newRes := append([]int{},currentRet...)
		*ret=append(*ret, newRes)
	}
	visited[index]=
}
