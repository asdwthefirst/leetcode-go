/**
  @author: jiangxi
  @since: 2022/12/28
  @desc: //TODO
**/
package divideandconquer

import "strconv"

func diffWaysToCompute(expression string) []int {
	ret := []int{}
	for i := 0; i < len(expression); i++ {
		switch expression[i : i+1] {
		case "-":
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					ret = append(ret, l-r)
				}
			}
		case "+":
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					ret = append(ret, l+r)
				}
			}
		case "*":
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					ret = append(ret, l*r)
				}
			}
		}
	}
	if len(ret) == 0 {
		num, _ := strconv.ParseInt(expression, 10, 64)
		ret = append(ret, int(num))
	}
	return ret
}
