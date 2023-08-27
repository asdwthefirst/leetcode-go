/*
*

	@author: jiangxi
	@since: 2023/8/27
	@desc: //TODO

*
*/
package newadd

import "sort"

func nextPermutation(nums []int) {
	//如果不是倒顺序，
	//从后往前，后两者是正序，直接调换，
	//假设倒的i+1,...,n-1都是倒顺序，但是i<i+1,向后找最后一个比i大的k，将i，k交换，此时i位后是逆序，直接倒序即可
	n := len(nums)
	i := n - 1

	for i > 0 {
		if nums[i-1] < nums[i] {
			break
		}
		i = i - 1
	}
	if i == 0 {
		sort.Ints(nums)
		return
	}
	i = i - 1
	k := i + 1
	for k < n && nums[k] > nums[i] {
		k++
	}
	k--
	nums[i], nums[k] = nums[k], nums[i]
	for j := i + 1; j <= (i+n-1)/2; j++ {
		nums[j], nums[n-j+i] = nums[n-j+i], nums[j]
	}

	return
}
