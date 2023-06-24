/*
*

	@author: jiangxi
	@since: 2023/6/24
	@desc: //TODO

*
*/
package backtracking

import "strconv"

//给你一个正整数 n ，请你返回 n 的 惩罚数 。
//
//n 的 惩罚数 定义为所有满足以下条件 i 的数的平方和：
//
//1 <= i <= n
//i * i 的十进制表示的字符串可以分割成若干连续子字符串，且这些子字符串对应的整数值之和等于 i 。

var sumArr [1001]int

func init() {
	sumArr[0] = 0
	for i := 1; i <= 1000; i++ {
		sumArr[i] = sumArr[i-1]
		if backtracking(strconv.Itoa(i*i), 0, 0, i) {
			sumArr[i] += i * i
		}
	}
}

func punishmentNumber(n int) int {
	return sumArr[n]
}

func backtracking(str string, pos int, sum, target int) bool {
	x := 0
	if pos == len(str) {
		return sum == target
	}
	for i := pos; i < len(str); i++ {
		x = x*10 + int(str[i]-'0')
		if backtracking(str, i+1, sum+x, target) {
			return true
		}
	}
	return false
}
