/*
*

	@author: jiangxi
	@since: 2022/12/20
	@desc: //TODO

*
*/
package sort

import (
	"math/rand"
	"time"
)

//超时，不知道优化点在哪
/*
func FindKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := quickSelect(nums, 0, len(nums)-1, k)
	fmt.Println(result)
	fmt.Println(time.Since(start))
	return result
}

func quickSelect(nums []int, left, right int, k int) int {
	if left >= right {
		return nums[left]
	}
	pa, pb := left, right
	randPos := rand.Int()%(right-left+1) + left
	nums[left], nums[randPos] = nums[randPos], nums[left]
	mid := nums[left]
	for pa < pb {
		for pa < pb && nums[pb] < mid {
			pb--
		}
		nums[pa] = nums[pb]
		for pa < pb && nums[pa] > mid {
			pa++
		}
		nums[pb] = nums[pa]
	}

	nums[pa] = mid
	//x:=nums[right]
	//i:=left-1
	//for j:=left;j<right;i++{
	//	if nums[j]<=x{}
	//}
	//
	fmt.Println(nums)
	fmt.Println(pa, pb, k)
	if pa == k-1+left {
		return nums[pa]
	}
	if pa < k-1+left {
		return quickSelect(nums, pa+1, right, k-pa-1+left)
	} else {
		return quickSelect(nums, left, pa-1, k)
	}
}
*/
//0624这是gpt优化的，把递归改为迭代之后能通过，之后再看把
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	// start := time.Now()
	result := quickSelect(nums, k)
	// fmt.Println(result)
	// fmt.Println(time.Since(start))
	return result
}

func quickSelect(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		randPos := rand.Int()%(right-left+1) + left
		nums[left], nums[randPos] = nums[randPos], nums[left]
		mid := nums[left]
		pa, pb := left, right
		for pa < pb {
			for pa < pb && nums[pb] < mid {
				pb--
			}
			nums[pa] = nums[pb]
			for pa < pb && nums[pa] >= mid {
				pa++
			}
			nums[pb] = nums[pa]
		}
		nums[pa] = mid
		if pa == k-1 {
			return nums[pa]
		} else if pa < k-1 {
			left = pa + 1
		} else {
			right = pa - 1
		}
	}
	return -1
}

/*
func findKthLargest(nums []int, k int) int {
    start:=time.Now()
    rand.Seed(time.Now().UnixNano())
    result:=quickSelect(nums, 0, len(nums)-1, len(nums)-k)
    fmt.Println(time.Since(start))
    return result

}

func quickSelect(a []int, l, r, index int) int {
    q := randomPartition(a, l, r)
    if q == index {
        return a[q]
    } else if q < index {
        return quickSelect(a, q + 1, r, index)
    }
    return quickSelect(a, l, q - 1, index)
}

func randomPartition(a []int, l, r int) int {
    i := rand.Int() % (r - l + 1) + l
    a[i], a[r] = a[r], a[i]
    return partition(a, l, r)
}

func partition(a []int, l, r int) int {
    x := a[r]
    i := l - 1
    for j := l; j < r; j++ {
        if a[j] <= x {
            i++
            a[i], a[j] = a[j], a[i]
        }
    }
    a[i+1], a[r] = a[r], a[i+1]
    return i + 1
}
*/

/*
func findKthLargest(nums []int, k int) int {
	buildMaxHeap(nums, len(nums))
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums, 0, i)
	}
	return nums[0]
}

func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize/2 - 1; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
}

func maxHeapify(nums []int, i, heapSize int) {
	left, right, largest := i*2+1, i*2+2, i
	if left < heapSize && nums[left] > nums[largest] {
		largest = left
	}
	if right < heapSize && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapify(nums, largest, heapSize)
	}

}*/
