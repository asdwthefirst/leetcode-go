/**
  @author: jiangxi
  @since: 2022/12/27
  @desc: //TODO
**/
package greed

//跟121是完全不同的思路，，把买入卖出时机划分到最小，只要有收益就计入，我还是没有想到
func maxProfit(prices []int) (ret int) {
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			ret += prices[i] - prices[i-1]
		}

	}
	return
}
