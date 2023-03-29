/**
  @author: jiangxi
  @since: 2022/11/11
  @desc: //TODO
**/
package stackqueue

func nextGreaterElements(nums []int) []int {
	var stack []int
	res := make([]int, len(nums))
	for i := range res {
		res[i] = -1
	}
	for i := 0; i < len(nums)*2; i++ {
		realIndex := i % len(nums)
		//fmt.Println(realIndex)
		//fmt.Println(stack)
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[realIndex] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[index] = nums[realIndex]
		}
		stack = append(stack, realIndex)
	}
	return res

}
