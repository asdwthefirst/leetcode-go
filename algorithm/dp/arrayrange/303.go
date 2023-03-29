/**
  @author: jiangxi
  @since: 2023/1/3
  @desc: //TODO
**/
package arrayrange

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]

	}
	numArray := NumArray{
		sums: sums,
	}
	return numArray
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > 0 {
		return this.sums[right] - this.sums[left-1]
	}
	return this.sums[right]

}
