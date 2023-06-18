/*
*

	@author: jiangxi
	@since: 2022/12/8
	@desc: //TODO

*
*/
package arraymatrix

func findErrorNums(nums []int) []int {
	res := make([]int, 2)
	for i, _ := range nums {
		//for nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
		//上面是原文的判断，下面是我的判定，而我的判定也确实通过了，
		//我判定是nums[nums[i]-1] == nums[i]才会陷入死循环，否则nums[nums[i]-1]就会得到正确的数
		//原文的判定，是如果nums[i]的位置已经正确了，那么也不应该去动了，也是合理的。
		for nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res[1] = i + 1
			res[0] = nums[i]
		}
	}
	return res
}

//	func findErrorNums(nums []int) []int {
//		result := make([]int, 2)
//		numMap := make(map[int]int)
//		for _, v := range nums {
//			numMap[v]++
//		}
//		for i := 1; i <= len(nums); i++ {
//			count, ok := numMap[i]
//			if !ok {
//				result[1] = i
//			}
//			if count == 2 {
//				result[0] = i
//			}
//		}
//		return result
//	}
//
//	func findErrorNums1(nums []int) []int {
//		quickSort(nums, 0, len(nums)-1)
//		fmt.Println(nums)
//		result := make([]int, 2)
//		for i := 1; i < len(nums); i++ {
//			if nums[i] == nums[i-1] {
//				fmt.Println(i)
//				result[0] = nums[i]
//			}
//			if nums[i]-nums[i-1] == 2 {
//				fmt.Println("dot2", i)
//				result[1] = nums[i] - 1
//			}
//		}
//		if nums[0] != 1 {
//			result[1] = 1
//		}
//		if nums[len(nums)-1] != len(nums) {
//			result[1] = len(nums)
//		}
//		return result
//	}
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	oldLeft, oldRight := left, right
	mid := nums[left]
	for left < right {
		for right > left && nums[right] >= mid {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= mid {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = mid
	quickSort(nums, oldLeft, left-1)
	quickSort(nums, left+1, oldRight)
}
