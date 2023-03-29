/**
  @author: jiangxi
  @since: 2022/11/11
  @desc: //TODO
**/
package stackqueue

//我想不到！单调栈
func dailyTemperatures(temperatures []int) []int {
	var stack []int
	res := make([]int, len(temperatures))
	for i := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[index] = i - index
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		index := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res[index] = 0
	}
	return res
}
