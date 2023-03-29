/**
  @author: jiangxi
  @since: 2022/12/8
  @desc: //TODO
**/
package arraymatrix

import "fmt"

func findErrorNums(nums []int) []int {
	result := make([]int, 2)
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v]++
	}
	for i := 1; i <= len(nums); i++ {
		count, ok := numMap[i]
		if !ok {
			result[1] = i
		}
		if count == 2 {
			result[0] = i
		}
	}
	return result
}

func findErrorNums1(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	result := make([]int, 2)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			fmt.Println(i)
			result[0] = nums[i]
		}
		if nums[i]-nums[i-1] == 2 {
			fmt.Println("dot2", i)
			result[1] = nums[i] - 1
		}
	}
	if nums[0] != 1 {
		result[1] = 1
	}
	if nums[len(nums)-1] != len(nums) {
		result[1] = len(nums)
	}
	return result
}

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
