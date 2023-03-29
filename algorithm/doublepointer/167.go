/**
  @author: jiangxi
  @since: 2022/12/19
  @desc: //TODO
**/
package doublepointer

func twoSum(numbers []int, target int) []int {
	pa, pb := 0, len(numbers)-1
	for pa < pb {
		if numbers[pa]+numbers[pb] == target {
			return []int{pa + 1, pb + 1}
		} else if numbers[pa]+numbers[pb] > target {
			pb--
		} else {
			pa++
		}
	}
	return nil

}
