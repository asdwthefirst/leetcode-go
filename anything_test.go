/**
  @author: jiangxi
  @since: 2022/10/31
  @desc: //TODO
**/
package leetcode_go

import (
	"fmt"
	"leetcode-go/algorithm/search/backtracking"
	"leetcode-go/algorithm/sort"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	fmt.Println(sort.FindKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

func TestPermuteUnique(t *testing.T) {
	backtracking.PermuteUnique([]int{-1, 2, -1, 2, 1, -1, 2, 1})
}
