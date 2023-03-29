/**
  @author: jiangxi
  @since: 2022/12/27
  @desc: //TODO
**/
package sort

import "fmt"

//首尾两个指针,两个指针的处理不太一样
func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	zero, one, two := 0, 0, len(nums)-1 //zero，one放同一个起点，这样可以每次换完之后不用重新判断而是直接前进，如果重新判断会导致进入0和0互换的循环
	for one <= two {
		switch nums[one] {
		case 0:
			fmt.Println(0, zero, one, two)
			nums[one], nums[zero] = nums[zero], nums[one]
			one++
			zero++
		case 1:
			fmt.Println(1, zero, one, two)
			one++
		case 2:
			fmt.Println(2, zero, one, two)
			nums[one], nums[two] = nums[two], nums[one]
			two-- //要重新判断one的位置是什么，one不前进
		}
	}
	return
}
