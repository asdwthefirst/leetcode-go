/**
  @author: jiangxi
  @since: 2022/11/11
  @desc: //TODO
**/
package hashset

func twoSum(nums []int, target int) []int {
	hashSet := make(map[int]int)
	for i := range nums {
		if _, ok := hashSet[target-nums[i]]; ok {
			return []int{i, hashSet[target-nums[i]]}
		}
		hashSet[nums[i]] = i
	}
	return []int{}

}
