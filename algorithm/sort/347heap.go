/**
  @author: jiangxi
  @since: 2022/12/27
  @desc: //TODO
**/
package sort

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value}) //会进行排序
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(h).([2]int)[0]
	}
	return result
}

type IHeap [][2]int

func (I IHeap) Len() int {
	return len(I)
}

func (I IHeap) Less(i, j int) bool {
	return I[i][1] < I[j][1]
}

func (I IHeap) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}

func (I *IHeap) Push(x interface{}) {
	*I = append(*I, x.([2]int))
}

func (I *IHeap) Pop() interface{} {
	x := (*I)[len(*I)-1]
	*I = (*I)[:len(*I)-1]
	return x
}
