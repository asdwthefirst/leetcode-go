/**
  @author: jiangxi
  @since: 2022/12/27
  @desc: //TODO
**/
package greed

//我也没想到，真的很简单
func maxProfit121(prices []int) int {
	min := prices[0]
	ret := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > ret {
			ret = prices[i] - min
		}
	}
	return ret
}
