package merge_arrays

// merge : fro mhttps://leetcode.com/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	i := m - 1     // Último índice de elementos válidos em nums1
	j := n - 1     // Último índice de nums2
	k := m + n - 1 // Último índice total de nums1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	return nums1
}
