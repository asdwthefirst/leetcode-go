/**
  @author: jiangxi
  @since: 2022/12/31
  @desc: //TODO
**/
package fibonacci

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	i1, i2 := 1, 2
	for i := 2; i < n; i++ {
		i1, i2 = i2, i1+i2
	}
	return i2

}
