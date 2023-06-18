/*
*

	@author: jiangxi
	@since: 2022/11/11
	@desc: //TODO

*
*/
package stackqueue

//给定一个整数数组temperatures， 表示每天的温度，
//返回一个数组answer，其中answer[i]是指对于第 i 天，下一个更高温度出现在几天后。
//如果气温在这之后都不会升高，请在该位置用0 来代替。

// 我想不到！单调栈

// 单调栈思想：
func dailyTemperatures(temperatures []int) []int {
	var stack []int
	res := make([]int, len(temperatures))
	for i := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1] //把暂时还没找到下家的成员存在stack中，可以写入时将成员pop，是个单调减栈
			res[index] = i - index
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		index := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res[index] = 0 //这个似乎是多余的，因为默认值0
	}
	return res
}
