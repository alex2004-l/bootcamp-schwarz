func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	left, right := 0, 0
	result := [][]int{}

	for left < len(nums1) && right < len(nums2) {
		if nums1[left][0] == nums2[right][0] {
			result = append(result, []int{nums1[left][0], nums1[left][1] + nums2[right][1]})
			left++
			right++
		} else if nums1[left][0] < nums2[right][0] {
			result = append(result, nums1[left])
			left++
		} else {
			result = append(result, nums2[right])
			right++
		}
	}

	for left < len(nums1) {
		result = append(result, nums1[left])
		left++
	}

	for right < len(nums2) {
		result = append(result, nums2[right])
		right++
	}

	return result
}