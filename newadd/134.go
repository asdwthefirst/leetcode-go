/*
*

	@author: jiangxi
	@since: 2023/7/28
	@desc: //TODO

*
*/
package newadd

import "fmt"

//如果能从x，到达y，但是到达不了y+1，选xy之间z，x能到达z，则到z，油量是有补充的，
//则从z出发不可能到y+1。所以可以直接从y出发。
//
func canCompleteCircuit(gas []int, cost []int) int {
	resCnt := 0
	start := 0
	res := 0
	n := len(gas)
	for start < len(gas) {
		fmt.Println(start)
		if gas[start] == 0 {
			start += 1
			continue
		}
		cur := 0
		flag := true
		for i := 0; i < len(gas); i++ {
			// fmt.Println(start,i,cur+gas[(start+i)%n]-cost[(start+i)%n])
			cur += gas[(start+i)%n] - cost[(start+i)%n]
			// fmt.Println(start,i,cur)
			if cur < 0 {
				start = max134(start+i, start+1)
				flag = false
				break
			}
		}
		if flag {
			// fmt.Println(start,resCnt)
			resCnt++
			res = start
			if resCnt > 1 {
				return -1
			}
			start = start + 1
		}
	}
	if resCnt == 0 {
		return -1
	}
	return res
}

func max134(a, b int) int {
	if a > b {
		return a
	}
	return b
}
