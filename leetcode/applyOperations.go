func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = 2 * nums[i]
			nums[i+1] = 0
		}
	}

	idx := 0
	cnt_0 := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cnt_0 += 1
		} else {
			nums[idx] = nums[i]
			idx++
		}
	}

	for idx < len(nums) {
		nums[idx] = 0
		idx += 1
	}

	return nums
}