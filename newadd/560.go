/*
*

	@author: jiangxi
	@since: 2023/7/27
	@desc: //TODO

*
*/
package newadd

//一个sum数组，保存0-k的和
//sum[b]-sum[a]为[a+1,b]区间和，暴力遍历则时间复杂度n~2

//但是下面是都是正数的情况，存在负数没有想到怎么样进行规模缩减。
// 如果[a,b]区间和大于k，则不需要再扩张
// start从0开始，向右移动end，
//  当<target，end右移动
//  当==target，res+1，且start右移动，
//  当>target时，start右移动，

//加上hash表，空间换时间

//并且一次遍历完成，所以map里存的都是过往的前缀和，不需要考虑位置.
//但是要记得初始化0也算一个位置

func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	prefixSum := 0
	sumMap := make(map[int]int)
	sumMap[0] = 1 //这个地方忘记处理了。
	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		res += sumMap[prefixSum-k]
		sumMap[prefixSum] += 1
	}
	return res
}
