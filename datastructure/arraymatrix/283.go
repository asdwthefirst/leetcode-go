/**
  @author: jiangxi
  @since: 2022/11/12
  @desc: //TODO
**/
package arraymatrix

func moveZeroes(nums []int) {
	//i,j:=0,0
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		} else {
			continue
		}

	}
	for j < len(nums) {
		nums[j] = 0
		j++
	}
	return

}
