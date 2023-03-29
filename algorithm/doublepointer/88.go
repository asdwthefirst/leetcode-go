/**
  @author: jiangxi
  @since: 2022/12/19
  @desc: //TODO
**/
package doublepointer

func merge(nums1 []int, m int, nums2 []int, n int) {
	pa, pb := m-1, n-1
	for pos := m + n - 1; pos >= 0; pos-- {
		if pb < 0 || (pa >= 0 && nums1[pa] > nums2[pb]) {
			nums1[pos] = nums1[pa]
			pa--
		} else {
			nums1[pos] = nums2[pb]
			pb--
		}
	}
	return

}
