/**
  @author: jiangxi
  @since: 2022/11/11
  @desc: //TODO
**/
package hashset

func containsDuplicate(nums []int) bool {
	dup := make(map[int]bool)
	for _, num := range nums {
		if _, ok := dup[num]; ok {
			return true
		}
		dup[num] = true
	}
	return false

}
